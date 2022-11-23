// Timeouts are important for programs that connect to external
// resources or that otherwise need to bound execution time.

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	// wait for 1 second if the c1 get something in that time
	// then process else move on with a timeout
	// select are blocking in nature, they wait until once of the case
	// is meet
	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout 2")
	}
}

// Timeout 1
// two
