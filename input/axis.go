package input

import (
	"fmt"
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
		fmt.Println(ae.RawValue)
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

// HandleAxisMessage func
func HandleAxisMessage(e Event) byte {
	ae := AxisEvent{axisMapping[e.Code], e.Value}
	_ = ae

	return 1
}
