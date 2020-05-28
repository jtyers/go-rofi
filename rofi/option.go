package rofi

import (
	"io"
)

// Option is an interface representing an option for a RofiFactory.
type Option interface {
	// ApplyCmdArgs supplies the arguments that implement this Option.
	//
	// The existing arguments, configured by other options in the Factory, are passed in,
	// which may affect the Arguments returned. These can be ignored if no validation is
	// needed. Note that only *new* arguments should be returned by this function.
	//
	// An error should be returned if, for example, the Arguments this Option would
	// apply are somehow incompatible with the supplied input arguments.
	ProvideArguments(existingArgs Arguments) (Arguments, error)
}

// Option is an interface providing stdin for a RofiFactory.
type StdinOption interface {
	Option // make this a 'sub interface' of Option

	// ProvideStdin provides a standard input stream to Rofi. If this option does not
	// need to provide a stdin stream, this function should return nil.
	ProvideStdin() io.Reader
}
