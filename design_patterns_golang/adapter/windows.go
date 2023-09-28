package main

import "fmt"

// actual window machine which have different method then
// what is called by client while plugging connector
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// create an adaptor so that client can call it's usual method
// it know about and adaptor handel the conversion form client method
// call to actual method call which is supported by windows machine

// this way adaptor satisfy the computer interface but actual
// windows computer do not
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
