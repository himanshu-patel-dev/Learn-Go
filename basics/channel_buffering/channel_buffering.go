package main

import "fmt"

func main() {
	// declaring a channel of type string with
	// capacity of max 2 items at any time
	channel := make(chan string, 2)
	channel <- "first"
	channel <- "second"

	// if we add third input before consuming anything
	// then it makes main goroutine wait for forever - deadlock
	// channel <- "third"

	// 	fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [chan send]:
	// main.main()
	//         /home/himanshu/HP/dev/Learn-Go/basics/channel_buffering/channel_buffering.go:14 +0x5c
	// exit status 2

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
