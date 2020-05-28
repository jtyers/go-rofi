package rofi

// generate slice helpers for []string (requires github.com/jtyers/slice to be installed)
//go:generate slice -type string -dir stringslice

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/jtyers/rofi/stringslice"
)

// Arguments represents the arguments to pass to rofi, and is simply a string slice.
type Arguments []string

func (a Arguments) append(args Arguments) Arguments {
	us := []string(a)
	them := []string(args)

	return Arguments(stringslice.Concat(us, them))
}

func NewArguments(args ...string) Arguments {
	return Arguments(args)
}

// Factory creates Rofi instances. Use this to configure how Rofi is called and displayed.
type Factory struct {
	// The Options passed to the factory
	options []Option

	// The stdin passed to the Rofi command
	stdin io.Reader
}

func NewFactory() *Factory {
	return &Factory{}
}

// NewRofi creates a new Rofi instance based on the arguments configured in this Factory.
func (f *Factory) NewRofi() (*Rofi, error) {
	args := Arguments{}

	for _, option := range f.options {
		newArgs, err := option.ProvideArguments(args)
		if err != nil {
			return nil, fmt.Errorf("error applying %#v: %v", option, err)
		}

		args = append(args, newArgs)
	}

	cmd := exec.Command("rofi", args...)

	return &Rofi{cmd}, nil
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

// WithMode sets the Rofi mode (via -show <mode>). Constants are provided for the known rofi modes.
func (f *Factory) WithMode(mode string) *Factory {
	return f.withOption(&ModeOption{mode})
}

// WithDmenu puts rofi in Dmenu mode (via -dmenu), where stdin is presented to the user and selected items are output to stdout.
func (f *Factory) WithDmenu() *Factory {
	return f.withOption(&DmenuOption{})
}

// WithDmenu puts rofi in Dmenu mode (via -dmenu), where stdin is presented to the user and selected items are output to stdout.
func (f *Factory) WithStdin(stdin io.Reader) *Factory {
	return f.withOption(&StdinOption{})
}

func (f *Factory) WithMessage() *Factory {
	return f.withOption(&DmenuOption{})
}

// NewRofi is syntactic sugar for `NewFactory().WithOption(option).NewRofi()`
func NewRofi(options ...Option) (*Rofi, error) {
	f := Factory{options}

	return f.NewRofi()
}
