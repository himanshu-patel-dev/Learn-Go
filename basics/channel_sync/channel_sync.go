package main

import (
	"fmt"
	"time"
)

// The signal channel will be used to notify another goroutine
// that this function’s work is done.
func worker(signal chan bool) {
	fmt.Println("First")
	time.Sleep(time.Second)
	fmt.Println("Second")
	// if this go-rtn do not send any value then a
	// deadlock is detected and program quits
	// fatal error: all goroutines are asleep - deadlock!
	signal <- true
}

// We can use channels to synchronize execution across goroutines.
// Example of using a blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish, you may prefer to use
// a WaitGroup.

func main() {
	done := make(chan bool)
	go worker(done)
	// Block until we receive a notification from the worker on the channel.
	<-done // if we remove this line then main exit without waiting for another go-rtn
}

// First
// Second
