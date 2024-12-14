package input

// Button Event Codes
const (
	BtnShiftLeft = iota
	BtnShiftRight
)

var buttonMapping = map[byte]uint16{
	BtnShiftRight: 292,
	BtnShiftLeft:  293,
}

var prevButtonEvent *Event

// HandleButtonMessage func
func HandleButtonMessage(e Event) (shouldSend bool, newByte byte) {
	shouldSend = true

	if prevButtonEvent != nil &&
		prevButtonEvent.Type == e.Type &&
		prevButtonEvent.Code == e.Code &&
		prevButtonEvent.Value == 1 && e.Value == 0 {
		switch e.Code {
		case buttonMapping[BtnShiftRight]:
			newByte = 1
		case buttonMapping[BtnShiftLeft]:
			newByte = 2

		default:
			shouldSend = false
			newByte = 0 // Unknown button event
		}
	}

	prevButtonEvent = &e
	return shouldSend, newByte
}
