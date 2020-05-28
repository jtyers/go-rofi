package rofi

import (
	"bytes"
	"os/exec"
	"strings"
)

// Rofi is a running Rofi instance created by a RofiFactory.
type Rofi struct {
	// The rofi process
	cmd *exec.Cmd
}

// Run executes the Rofi command. Stdout and Stderr are returned, followed by an error if encountered.
func (r *Rofi) Run(stdin string) (string, string, error) {
	r.cmd.Stdin = strings.NewReader(stdin)

	stdout := bytes.Buffer{}
	r.cmd.Stdout = &stdout

	stderr := bytes.Buffer{}
	r.cmd.Stderr = &stderr

	err := r.cmd.Run()

	return stdout.String(), stderr.String(), err
}
