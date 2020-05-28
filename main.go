package main

import (
	"fmt"

	"github.com/jtyers/go-rofi/rofi"
)

func main() {
	r, err := rofi.NewFactory().WithDmenu("first", "second", "third").WithMessage("you must choose one, Luke").NewRofi()
	result, err := r.Run()
	fmt.Printf("result: %#v\nerr: %v\n", result, err)
}
