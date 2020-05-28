package rofi

// MessageOption puts Rofi in dmenu mode, displaying stdin and outputting selected options to stdout.
type MessageOption struct {
	message string
}

func (o *MessageOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-mesg", o.message), nil
}

func (f *Factory) WithMessage(message string) *Factory {
	return f.withOption(&MessageOption{message})
}
