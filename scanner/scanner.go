package scanner

import (
	"log"
	"time"

	"go.bug.st/serial"
)

func StartScanner(portName string, baud int, onScan func(code string)) {
	mode := &serial.Mode{
		BaudRate: baud,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Fatalf("Failed to open serial port: %v", err)
	}

	port.SetReadTimeout(20 * time.Millisecond)

	log.Println("Barcode scanner connected:", portName)

}
