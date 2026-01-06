package devices

import (
	"fmt"
	"log"
	"time"

	"gitea.locker98.com/locker98/Store-Manager-Pro/routes"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

func Devices() {
	// create variables for device port address and barcode data
	var device string
	data := make([]byte, 256)

	for {
		// Check if a device has been found
		if device == "" {
			// If no device found search for a new device for a maximum of 10 seconds
			device = FindDevice(10)
		} else {
			// Once a device is found wait for it to get fully connected to the computer and setup
			time.Sleep(1000 * time.Millisecond)

			// Try to open the device port
			port, err := OpenScanner(device, 115200)

			// Restart scanning if their is an error
			if err != nil {
				fmt.Println("Opening Error:", err)
				device = ""
				continue
			}

			// Pull device data
			var barcode string
			for {
				// Read data from the scanner
				n, err := port.Read(data)
				// Check for errors
				if err != nil {
					device = ""
					fmt.Println("Read Error", err)
					break
				}

				// Check if the barcode scanner returned more of the barcode or if it is just an empty string
				if string(data[:n]) == "" && barcode != "" {
					// Clear Send the barcode now that their is no more data and clear the barcode
					routes.SendWSBarcodes(barcode)
					barcode = ""
				} else {
					// append the chunk of the barcode to the current value
					barcode += string(data[:n])
				}
			}
		}
	}
}

func OpenScanner(portName string, baud int) (serial.Port, error) {
	mode := &serial.Mode{
		BaudRate: baud,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		fmt.Printf("Failed to open serial port: %v\n", err)
		return nil, err
	}

	port.SetReadTimeout(100 * time.Millisecond)
	return port, nil
}

func FindDevice(timeout int) string {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	timeoutTimer := time.NewTimer(time.Duration(timeout) * time.Second)
	defer timeoutTimer.Stop()

	// Track seen ports and when they appeared
	seen := make(map[string]time.Time)

	for {
		select {
		case <-ticker.C:
			ports, err := enumerator.GetDetailedPortsList()
			if err != nil {
				log.Println("serial scan error:", err)
				continue
			}

			now := time.Now()
			current := make(map[string]bool)

			for _, port := range ports {
				current[port.Name] = true

				// New device appeared
				if _, ok := seen[port.Name]; !ok {
					seen[port.Name] = now
				}
			}

			// Check for devices that disappeared quickly
			for name, firstSeen := range seen {
				if !current[name] {
					if now.Sub(firstSeen) <= 2*time.Second {
						// Device appeared & disappeared quickly → barcode scanner behavior
						return name
					}
					delete(seen, name)
				}
			}

		case <-timeoutTimer.C:
			return ""
		}
	}
}
