package main

import (
	"fmt"
	"os"
)

// go run hello.go uses run and hello.go arguments to the go program

func main() {
	// os.Args provides access to raw command-line arguments.
	// Note that the first value in this slice is the path to
	// the program, and os.Args[1:] holds the arguments to the program
	argsWithProg := os.Args
	fmt.Println("File path: ", argsWithProg[0])
	fmt.Println("CMD Line Parameters: ", argsWithProg[1:])
}

// $ go run cmd/cmd.go a b c d
// File path:  /tmp/go-build607469594/b001/exe/cmd
// CMD Line Parameters:  [a b c d]
