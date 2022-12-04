// A line filter is a common type of program that reads input on stdin,
// processes it, and then prints some derived result to stdout.
// grep and sed are common line filters.

// writes a capitalized version of all input text
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Wrapping the unbuffered os.Stdin with a buffered scanner
	// gives us a convenient Scan method that advances the scanner
	// to the next token; which is the next line in the default scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Text returns the current token, here the next line, from the input.
		fmt.Println(strings.ToUpper(scanner.Text()))
	}

	// scanner.Scan return false when it stops either due to end of input
	// or face an error, so it's good to check for any error when Scan stops
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// -- from console input
// Hello
// HELLO
// there
// THERE
// ^Csignal: interrupt

// -- from file
// $ cat data.txt  | go run line_filter.go
// FIRST
// SECOND
