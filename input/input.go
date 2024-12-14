// Package input is responsible for listening to input events and parsing them
package input

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unsafe"
)

var devicePath string

// TimeVal struct
type TimeVal struct {
	Sec  uint64
	USec uint64
}

// Event struct
type Event struct {
	Time  TimeVal
	Type  uint16
	Code  uint16
	Value uint32
}

// Event Codes
const (
	EventSyn = 0x00
	EventKey = 0x01
	EventAbs = 0x03
	EventMsc = 0x04
)

func init() {
	devices, err := filepath.Glob("/dev/input/event*")
	if err != nil {
		fmt.Println("Unable to list events")
		os.Exit(1)
	}

	if len(devices) == 0 {
		fmt.Println("No devices found")
		os.Exit(1)
	}

	for _, device := range devices {
		deviceNamePath := fmt.Sprintf("/sys/class/input/%s/device/name", filepath.Base(device))
		nameBytes, err := os.ReadFile(deviceNamePath)
		if err != nil {
			fmt.Println("Unable to read device name for event:", device)
			continue
		}

		name := string(nameBytes)

		if strings.Index(name, "Logitech G923") != -1 {
			fmt.Println("FOUND ->", name)
			devicePath = device
		}
	}

	if devicePath == "" {
		fmt.Println("Unable to find the device")
		os.Exit(1)
	}
}

// ListenToEvents func
func ListenToEvents(ch chan<- byte) {
	fmt.Println("Racing Wheel is at:", devicePath)
	f, err := os.Open(devicePath)
	if err != nil {
		fmt.Println("Unable to read the file:", devicePath)
		os.Exit(1)
	}
	defer func() {
		f.Close()
		close(ch)
	}()

	event := Event{}
	size := int(unsafe.Sizeof(event))
	buffer := make([]byte, size)

	fmt.Println("Listening To Events...")

	for {
		_, err := io.ReadFull(f, buffer)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("Error reading the event: %v\n", err)
			continue
		}

		event.Type = binary.LittleEndian.Uint16(buffer[16:18])
		event.Code = binary.LittleEndian.Uint16(buffer[18:20])
		event.Value = binary.LittleEndian.Uint32(buffer[20:24])

		if event.Type == EventSyn || event.Type == EventMsc {
			continue
		}

		var newByte byte
		var shouldSend bool
		switch event.Type {
		case EventKey:
			shouldSend, newByte = HandleButtonMessage(event)

		case EventAbs:
			shouldSend, newByte = HandleAxisMessage(event)
		}

		if shouldSend {
			ch <- newByte
		}
	}
}
