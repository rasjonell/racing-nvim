package input

import "fmt"

var prevButtonEvent *Event

// HandleButtonMessage func
func HandleButtonMessage(e Event) byte {
	if prevButtonEvent != nil &&
		prevButtonEvent.Type == e.Type &&
		prevButtonEvent.Code == e.Code &&
		prevButtonEvent.Value == 1 && e.Value == 0 {
		fmt.Println("Clicked")
	}

	prevButtonEvent = &e
	return 0
}
