package main

import "fmt"

// interface of printer
type Printer interface {
	PrintFile()
}

// epson concrete implementation
type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// HP concrete implementation
type HP struct{}

func (p *HP) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
