package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and
// receive those values into another goroutine.

func main() {
	// Create a new channel with make(chan val-type).
	// Channels are typed by the values they convey
	message := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	go func() { message <- "ping" }()

	// The <-channel syntax receives a value from the channel.
	msg := <-message
	fmt.Println(msg) // ping

	// By default sends and receives block until both the sender and receiver
	// are ready. This property allowed us to wait at the end of our program
	// for the "ping" message without having to use any other synchronization.
}
