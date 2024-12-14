package input

import (
	"math"
)

// AxisEvent struct
type AxisEvent struct {
	AxisType byte
	RawValue uint32
}

// Axis Event Codes
const (
	AxisWheel = iota
	AxisBreak
	AxisClutch
	AxisAccelerate
	AxisDPadVertical
	AxisDPadHorizontal
)

var axisMapping = map[uint16]byte{
	0:  AxisWheel,
	5:  AxisBreak,
	1:  AxisClutch,
	2:  AxisAccelerate,
	17: AxisDPadVertical,
	16: AxisDPadHorizontal,
}

// ParsedValue func
func (ae *AxisEvent) ParsedValue() uint32 {
	switch ae.AxisType {
	case AxisWheel:
		return uint32((float64(ae.RawValue) / 65535.0 * 900))

	case AxisBreak:
		return uint32(math.Abs(float64(255 - ae.RawValue)))

	case AxisClutch:
		return uint32(math.Abs(float64(255 - ae.RawValue)))

	case AxisAccelerate:
		return uint32(math.Abs(float64(255 - ae.RawValue)))

	default:
		return ae.RawValue
	}
}

var prevWheelValue uint32
var distanceThreshold uint32 = 10

// HandleAxisMessage func
func HandleAxisMessage(e Event) (shouldSend bool, newByte byte) {
	ae := AxisEvent{axisMapping[e.Code], e.Value}

	switch ae.AxisType {
	case AxisWheel:
		parsed := ae.ParsedValue()
		distance := uint32(math.Abs(float64(prevWheelValue) - float64(parsed)))
		if distance >= distanceThreshold {
			newByte = 3 // Decrease The Char
			if parsed > prevWheelValue {
				newByte = 4 // Increase The Char
			}

			prevWheelValue = ae.ParsedValue()

			return true, newByte
		}

	default:
		newByte = 0 // Unknown Axis Event
	}

	return false, 0
}
