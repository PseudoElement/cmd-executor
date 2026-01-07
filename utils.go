package main

import "fmt"

func assert(ok bool, panicMsg string, err error) {
	if !ok {
		var msg error
		if err != nil {
			msg = fmt.Errorf("assert panicked: %s. Original error: %w.", panicMsg, err)
		} else {
			msg = fmt.Errorf("assert panicked: %s.", panicMsg)
		}
		panic(msg)
	}
}

func logGreen(msg string) {
	green := "\033[32m"
	reset := "\x1b[0m"
	fmt.Printf("%s%s%s\n", green, msg, reset)
}

func logRed(msg string) {
	red := "\033[31m"
	reset := "\x1b[0m"
	fmt.Printf("%s%s%s\n", red, msg, reset)
}

func logBlue(msg string) {
	blue := "\033[34m"
	reset := "\x1b[0m"
	fmt.Printf("%s%s%s\n", blue, msg, reset)
}
