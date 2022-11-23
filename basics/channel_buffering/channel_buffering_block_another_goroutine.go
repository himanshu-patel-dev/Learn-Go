package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 2)
	channel <- "first"
	channel <- "second"

	// if we add third input before consuming anything
	// then it makes main goroutine wait for forever - deadlock
	go func(ch chan string) {
		fmt.Println("Put Third") // 1
		ch <- "third"            // third insert will make this go-rtn to wait infinitely
	}(channel)

	// so that the main routine wait and don't exist before 1st print
	time.Sleep(1 * time.Second)
	fmt.Println("Done") // 2
}

// Put Third
// Done

// This happens because main don't know existence of other goroutines and
// continue and end up exiting before the another goroutine stuck in deadlock
// can continue
