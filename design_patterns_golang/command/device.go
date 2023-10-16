package main

import "fmt"

// Device Interface
type Device interface {
	on()
	off()
}

// Concrete Implementation

type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("TV is On")
}

func (t *TV) off() {
	t.isRunning = false
	fmt.Println("TV is Off")
}
