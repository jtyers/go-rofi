package rofi

import "fmt"

// ColumnsOption sets the number of columns to display (i.e. width)
type ColumnsOption struct {
	columns int
}

func (o *ColumnsOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-columns", fmt.Sprintf("%d", o.columns)), nil
}

func (f *Factory) WithColumns(columns int) *Factory {
	return f.withOption(&ColumnsOption{columns})
}
