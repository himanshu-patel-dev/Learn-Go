package main

import "fmt"

/*

Bridge is a structural design pattern that divides business logic
or huge class into separate class hierarchies that can be developed
independently.

*/

func main() {
	// here instead of 4 class 2 (computer) x 2 (printer)
	// we just managed to do this with 2 + 2 classes
	// this becomes more resonating when the n * m is too
	// large then n + m
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
