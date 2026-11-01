package devices

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gitea.locker98.com/locker98/Store-Manager-Pro/routes"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

const (
	zebraVendorID = "05e0"

	scanInterval    = 1 * time.Second
	readTimeout     = 100 * time.Millisecond
	barcodeIdleTime = 250 * time.Millisecond
)

func Devices() {
	for {
		// Automatically find a connected Zebra scanner.
		device := FindDevice()

		if device == "" {
			continue
		}

		log.Printf("Found Zebra scanner on %s", device)

		// Open the scanner.
		port, err := OpenScanner(device, 9600)
		if err != nil {
			log.Printf("Failed to open scanner %s: %v", device, err)
			time.Sleep(time.Second)
			continue
		}

		log.Printf("Connected to Zebra scanner on %s", device)

		// Read barcodes until the scanner disconnects.
		ReadScanner(port)

		port.Close()

		log.Println("Zebra scanner disconnected. Searching for scanner...")
	}
}

// FindDevice searches for a Zebra USB serial device.
// It also finds devices that were already connected at startup.
func FindDevice() string {
	ticker := time.NewTicker(scanInterval)
	defer ticker.Stop()

	for {
		ports, err := enumerator.GetDetailedPortsList()
		if err != nil {
			log.Printf("Serial device discovery error: %v", err)
		} else {
			for _, port := range ports {
				if IsZebraScanner(port) {
					return port.Name
				}
			}
		}

		<-ticker.C
	}
}

// IsZebraScanner identifies Zebra USB serial devices.
func IsZebraScanner(port *enumerator.PortDetails) bool {
	if port == nil || !port.IsUSB {
		return false
	}

	// Normalize VID: some platforms report it with a 0x prefix or odd casing.
	vid := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(port.VID)), "0x")
	if vid == zebraVendorID {
		return true
	}

	// Fall back to the product string for devices with a missing or odd VID.
	product := strings.ToLower(port.Product)
	return strings.Contains(product, "zebra") || strings.Contains(product, "symbol")
}

// OpenScanner opens the serial connection.
func OpenScanner(portName string, baud int) (serial.Port, error) {
	mode := &serial.Mode{
		BaudRate: baud,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		return nil, fmt.Errorf("open serial port %s: %w", portName, err)
	}

	if err := port.SetReadTimeout(readTimeout); err != nil {
		port.Close()
		return nil, fmt.Errorf("set serial timeout: %w", err)
	}

	return port, nil
}

// ReadScanner reads barcode data and sends completed scans.
func ReadScanner(port serial.Port) {
	data := make([]byte, 256)
	var barcode strings.Builder

	var lastData time.Time

	for {
		n, err := port.Read(data)

		if err != nil {
			log.Printf("Scanner read error: %v", err)
			return
		}

		if n > 0 {
			lastData = time.Now()

			for _, b := range data[:n] {
				switch b {
				case '\r', '\n':
					sendBarcode(&barcode)

				default:
					barcode.WriteByte(b)
				}
			}

			continue
		}

		// Some scanners don't send a line terminator.
		// Send the barcode after a short period of inactivity.
		if barcode.Len() > 0 &&
			!lastData.IsZero() &&
			time.Since(lastData) >= barcodeIdleTime {
			sendBarcode(&barcode)
		}
	}
}

// sendBarcode forwards a completed barcode.
func sendBarcode(barcode *strings.Builder) {
	value := strings.TrimSpace(barcode.String())
	barcode.Reset()

	if value == "" {
		return
	}

	log.Printf("Barcode scanned: %s", value)
	routes.SendWSBarcodes(value)
}
