package rofi

// DmenuOption puts Rofi in dmenu mode, displaying stdin and outputting selected options to stdout.
type DmenuOption struct {
}

func (o *DmenuOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-dmenu"), nil
}

func (o *DmenuOption) ProvideStdin() io.Reader {
	return nil
}
