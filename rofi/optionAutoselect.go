package rofi

// AutoselectOption sets the number of columns to display (i.e. autoselect)
type AutoselectOption struct {
}

func (o *AutoselectOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-autoselect"), nil
}

func (f *Factory) WithAutoselect() *Factory {
	return f.withOption(&AutoselectOption{})
}
