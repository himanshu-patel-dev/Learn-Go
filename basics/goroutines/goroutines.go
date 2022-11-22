package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("Direct")

	// this will continue in paraller and will not
	// block main thread from completion
	go f("goroutine")

	go func(s string) {
		fmt.Println(s)
	}("Going")
	// Our two function calls are running asynchronously
	// in separate goroutines now.

	// this make main thread wait for 1 sec
	time.Sleep(time.Second)
	fmt.Println("Done")
}

// Direct : 0
// Direct : 1
// Direct : 2
// Going
// goroutine : 0
// goroutine : 1
// goroutine : 2
// Done

// If we remove sleep time then main thread
// don't wait for other goroutines

// Direct : 0
// Direct : 1
// Direct : 2
// Done
