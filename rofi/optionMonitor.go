package rofi

import "fmt"

const (
	// MonitorWithFocus Show rofi on the currently focussed monitor.
	MonitorWithFocus = -1

	// MonitorAboveFocussedWindow Show rofi on top of the currently focussed window
	MonitorAboveFocussedWindow = -2

	// MonitorMousePosition Show rofi at the position of the mouse (still honours the location setting; use LocationTopLeft to get normal context menu behaviour)
	MonitorMousePosition = -3

	// MonitorWithFocussedWindow Show rofi on the monitor with the focused window
	MonitorWithFocussedWindow = -4

	// MonitorWithMousePointer Show rofi on monitor that contains the mouse pointer
	MonitorWithMousePointer = -5
)

// MonitorOption sets the number of columns to display (i.e. monitor)
type MonitorOption struct {
	monitor int
}

func (o *MonitorOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-monitor", fmt.Sprintf("%d", o.monitor)), nil
}

func (f *Factory) WithMonitor(monitor int) *Factory {
	return f.withOption(&MonitorOption{monitor})
}
