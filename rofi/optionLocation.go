package rofi

import "fmt"

const (
	LocationTopLeft      = 1
	LocationTopCenter    = 2
	LocationTopRight     = 3
	LocationMiddleLeft   = 8
	LocationMiddleCenter = 0
	LocationMiddleRight  = 4
	LocationBottomLeft   = 7
	LocationBottomCenter = 6
	LocationBottomRight  = 5
)

// LocationOption sets the number of columns to display (i.e. location)
type LocationOption struct {
	location int
}

func (o *LocationOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-location", fmt.Sprintf("%d", o.location)), nil
}

func (f *Factory) WithLocation(location int) *Factory {
	return f.withOption(&LocationOption{location})
}
