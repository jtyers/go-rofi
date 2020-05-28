package rofi

// RunModeOption puts Rofi in dmenu mode, displaying stdin and outputting selected options to stdout.
type ModeOption struct {
	mode string
}

func (o *ModeOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-show", o.mode), nil
}

// WithMode sets the Rofi mode (via -show <mode>). Constants are provided for the known rofi modes.
func (f *Factory) WithMode(mode string) *Factory {
	return f.withOption(&ModeOption{mode})
}
