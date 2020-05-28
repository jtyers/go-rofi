package Keystroke

type Keystroke int

// Map keystrokes to Keystroke instances
const (
	// Unknown - usually caused if user exits without entering input
	KeystrokeUnknown = Keystroke(1)

	KeystrokeEnter    = Keystroke(0)
	KeystrokeCustom1  = Keystroke(10)
	KeystrokeCustom2  = Keystroke(11)
	KeystrokeCustom3  = Keystroke(12)
	KeystrokeCustom4  = Keystroke(13)
	KeystrokeCustom5  = Keystroke(14)
	KeystrokeCustom6  = Keystroke(15)
	KeystrokeCustom7  = Keystroke(16)
	KeystrokeCustom8  = Keystroke(17)
	KeystrokeCustom9  = Keystroke(18)
	KeystrokeCustom10 = Keystroke(19)
)
