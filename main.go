// package main for running the input stuff
package main

import (
	"fmt"
	"wheeld/input"
)

func main() {
	ch := make(chan byte)

	go input.ListenToEvents(ch)

	for msg := range ch {
		fmt.Println("Got Msg:", msg)
	}
}
