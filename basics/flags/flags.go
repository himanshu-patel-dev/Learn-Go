package main

import (
	"flag"
	"fmt"
)

// Basic flag declarations are available for string, integer,
// and boolean options.

func main() {
	// flag name, default value, desc of flag
	// flag return a pointer to flag value and not flag value directly
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 41, "a number")
	boolPtr := flag.Bool("bool", false, "a boolean")

	var svar string
	flag.StringVar(&svar, "svar", "default value", "a string var")

	// Once all flags are declared, call flag.Parse() to execute
	// the command-line parsing.
	flag.Parse()

	fmt.Println(*wordPtr)
	fmt.Println(*numPtr)
	fmt.Println(*boolPtr)
	fmt.Println(svar)
	fmt.Println(flag.Args())
}

// $ go run flags.go
// foo
// 41
// false
// default value
// []

// $ go run flags.go -word=bar -numb=12 -bool=true -svar='updated value' a b c d
// bar
// 12
// true
// updated value
// [a b c d]
