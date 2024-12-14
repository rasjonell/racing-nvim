// package main for running the input stuff
package main

import (
	"fmt"
	"net"
	"os"
	"wheeld/input"
)

func main() {
	ch := make(chan byte)

	go input.ListenToEvents(ch)

	pipePath := os.Args[1]
	conn, err := net.Dial("unix", pipePath)
	if err != nil {
		fmt.Printf("Failed to connect to the socket: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	for msg := range ch {
		fmt.Println("Got Msg:", msg)
		_, err := conn.Write([]byte{msg})
		if err != nil {
			fmt.Printf("Failed to send a command: %v\n", err)
		}
	}
}
