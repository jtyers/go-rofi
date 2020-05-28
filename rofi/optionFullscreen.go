package rofi

// FullscreenOption sets the number of columns to display (i.e. fullscreen)
type FullscreenOption struct {
}

func (o *FullscreenOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-fullscreen"), nil
}

func (f *Factory) WithFullscreen() *Factory {
	return f.withOption(&FullscreenOption{})
}
