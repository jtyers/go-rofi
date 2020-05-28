package rofi

import "fmt"

// RowsOption sets the number of columns to display (i.e. width)
type RowsOption struct {
	rows int
}

func (o *RowsOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-rows", fmt.Sprintf("%d", o.rows)), nil
}

func (f *Factory) WithRows(rows int) *Factory {
	return f.withOption(&RowsOption{rows})
}
