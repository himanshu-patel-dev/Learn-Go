// Basic sends and receives on channels are blocking.
// However, we can use select with a default clause to implement
// non-blocking sends, receives, and even non-blocking multi-way selects.
package main

import "fmt"

func main() {
	message := make(chan string)
	signals := make(chan bool)

	// Here’s a non-blocking receive. If a value is available on messages then
	// select will take the <-messages case with that value. If not it will
	// immediately take the default case.
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("No message received")
	}

	// A non-blocking send works similarly. Here msg cannot be sent to the messages
	// channel, because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.
	msg := "Hello"
	select {
	case message <- msg:
		fmt.Println("Message feed into channel")
	default:
		fmt.Println("No Message feed into channel")
	}

	// We can use multiple cases above the default clause to implement a
	// multi-way non-blocking select. Here we attempt non-blocking receives on
	// both messages and signals.
	select {
	case msg := <-message:
		fmt.Println("Message Received", msg)
	case sig := <-signals:
		fmt.Println("Received Signal", sig)
	default:
		fmt.Println("No Activity")
	}
}

// No message received
// No Message feed into channel
// No Activity
