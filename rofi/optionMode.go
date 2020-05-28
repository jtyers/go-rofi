package rofi

import "io"

// RunModeOption puts Rofi in dmenu mode, displaying stdin and outputting selected options to stdout.
type ModeOption struct {
	mode string
}

func (o *ModeOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-show", o.mode), nil
}

func (o *ModeOption) ProvideStdin() io.Reader {
	return nil
}
