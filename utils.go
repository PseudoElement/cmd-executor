package main

import "fmt"

func assert(ok bool, panicMsg string) {
	if !ok {
		msg := fmt.Errorf("assert panicked: %s", panicMsg)
		panic(msg)
	}
}
