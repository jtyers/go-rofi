package rofi

// generate slice helpers for []Keystroke (requires github.com/jtyers/slice to be installed)
//go:generate slice -type Keystroke -dir slices -package slices -import github.com/jtyers/go-rofi/rofi/keystroke

import (
	"bytes"
	"io"
	"os/exec"
	"strings"

	k "github.com/jtyers/go-rofi/rofi/keystroke"
	"github.com/jtyers/go-rofi/rofi/slices"
)

func findKeystroke(exitCode int) k.Keystroke {
	filtered := slices.FilterKeystroke([]k.Keystroke{
		k.KeystrokeUnknown,
		k.KeystrokeEnter,
		k.KeystrokeCustom1,
		k.KeystrokeCustom2,
		k.KeystrokeCustom3,
		k.KeystrokeCustom4,
		k.KeystrokeCustom5,
		k.KeystrokeCustom6,
		k.KeystrokeCustom7,
		k.KeystrokeCustom8,
		k.KeystrokeCustom9,
		k.KeystrokeCustom10,
	}, func(k k.Keystroke, i int) bool {
		return int(k) == exitCode
	})

	if len(filtered) == 1 {
		return filtered[0]
	}

	return k.KeystrokeUnknown
}

// Rofi is a running Rofi instance created by a RofiFactory.
type Rofi struct {
	// The rofi process
	cmd *exec.Cmd

	// The rofi process
	stdin io.Reader
}

type RofiResult struct {
	Stdout []string
	Stderr string

	ExitCode int

	Keystroke k.Keystroke
}

// Run executes the Rofi command. Stdout and Stderr are returned, followed by an error if encountered.
func (r *Rofi) Run() (RofiResult, error) {
	r.cmd.Stdin = r.stdin

	stdout := bytes.Buffer{}
	r.cmd.Stdout = &stdout

	stderr := bytes.Buffer{}
	r.cmd.Stderr = &stderr

	err := r.cmd.Run()

	if _, ok := err.(*exec.ExitError); ok {
		err = nil // ignore bad exit codes
	}

	return RofiResult{
		Stdout: slices.NewStringSlice(strings.Split(stdout.String(), "\n")).Filter(func(x string, i int) bool {
			return x != ""
		}).Value(),
		Stderr:    stderr.String(),
		ExitCode:  r.cmd.ProcessState.ExitCode(),
		Keystroke: findKeystroke(r.cmd.ProcessState.ExitCode()),
	}, err
}
