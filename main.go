package main

import (
	"fmt"

	"github.com/jtyers/go-rofi/rofi"
)

func main() {
	r, err := rofi.NewFactory().WithMode(rofi.RofiModeWindow).NewRofi()
	stdout, stderr, err := r.Run("foo\nbar\nbaz")
	fmt.Printf("stdout: %q\nstderr: %q\nerr: %v\n", stdout, stderr, err)
}
