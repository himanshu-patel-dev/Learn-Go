package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// not extra cmd line arg given
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	// get cmd line args
	arguments := os.Args
	var err error = errors.New("An error")

	/*
		This is the trickiest part of the program because, if the first command-line argument is not a
		proper float, you will need to check the next one and keep checking until you find a suitable
		command-line argument. If none of the command-line arguments are in the correct format,
		errors.go will terminate and print a message on the screen. All this checking happens by
		examining the error value that is returned by strconv.ParseFloat() . All of this code is
		there just for the accurate initialization of the min and max variables.
	*/

	k := 1
	var n float64
	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, n
	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
