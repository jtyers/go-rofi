package rofi

import "fmt"

// WidthOption sets the number of columns to display (i.e. width)
type WidthOption struct {
	width int
}

func (o *WidthOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-width", fmt.Sprintf("%d", o.width)), nil
}

func (f *Factory) WithWidth(width int) *Factory {
	return f.withOption(&WidthOption{width})
}
