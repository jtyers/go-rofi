package rofi

// generate slice helpers for []string (requires github.com/jtyers/slice to be installed)
//go:generate slice -type string -dir slices -package slices

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/jtyers/go-rofi/rofi/slices"
)

// Arguments represents the arguments to pass to rofi, and is simply a string slice.
type Arguments []string

func (a Arguments) append(args Arguments) Arguments {
	us := []string(a)
	them := []string(args)

	return Arguments(slices.ConcatString(us, them))
}

func NewArguments(args ...string) Arguments {
	return Arguments(args)
}

// Factory creates Rofi instances. Use this to configure how Rofi is called and displayed.
type Factory struct {
	// The Options passed to the factory
	options []Option
}

func NewFactory() *Factory {
	return &Factory{}
}

// NewRofi creates a new Rofi instance based on the arguments configured in this Factory.
func (f *Factory) NewRofi() (*Rofi, error) {
	args := Arguments{}

	// The stdin passed to the Rofi command
	var stdin io.Reader

	for _, option := range f.options {
		newArgs, err := option.ProvideArguments(args)
		if err != nil {
			return nil, fmt.Errorf("error applying %#v: %v", option, err)
		}

		args = args.append(newArgs)

		if stdinOption, ok := option.(StdinOption); ok {
			// if this is also an stdinOption, then use it to provide stdin to Rofi,
			// but only if there is not already an stdin specified
			if stdin != nil {
				return nil, fmt.Errorf("error applying %#v: stdin already provided by a previous option", stdinOption)
			}

			stdin = stdinOption.ProvideStdin()
		}
	}

	cmd := exec.Command("rofi", args...)

	return &Rofi{cmd, stdin}, nil
}

// WithOption configures the given Option on this factory.
//
// The next call to NewRofi() will include the effects of the added option, but any existing instances of Rofi created by this
// factory are unaffected.

// Panics if an error occurs. This factory instance is returned, to allow for chained calls.
func (f *Factory) withOption(option Option) *Factory {
	f.options = append(f.options, option)
	return f
}

const (
	RofiModeRun      = "run"
	RofiModeDRun     = "drun"
	RofiModeWindow   = "window"
	RofiModeSsh      = "ssh"
	RofiModeCombi    = "combi"
	RofiModeKeys     = "keys"
	RofiModeWindowCd = "windowcd"
)

// NewRofi is syntactic sugar for `NewFactory().WithOption(option).NewRofi()`
func NewRofi(options ...Option) (*Rofi, error) {
	f := Factory{options: options}

	return f.NewRofi()
}
