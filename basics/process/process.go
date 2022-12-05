package main

import (
	"fmt"
	"os/exec"
)

// Sometimes our Go programs need to spawn other, non-Go processes.

func main() {
	// The exec.Command helper creates an object to
	// represent this external process.
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(">date")
	fmt.Println(string(dateOut))
	// >date
	// Mon Dec  5 09:27:56 AM IST 2022
}
