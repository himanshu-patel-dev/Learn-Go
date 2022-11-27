// Timers are for when you want to do something once in the future -
// tickers are for when you want to do something repeatedly at regular intervals.
package main

import (
	"fmt"
	"time"
)

func main() {
	// ticket which ticks after every 500 ms
	ticker := time.NewTicker(500 * time.Millisecond)
	// notify on done channel to end the for loop
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// main thread wait for 1600 ms to let goroutine work
	time.Sleep(1600 * time.Millisecond)
	// stopping the ticker after 1600 ms
	ticker.Stop()
	// send signal to ticker go routine to end
	done <- true
	fmt.Println("Ticker Stopped")
}

// Tick at 2022-11-26 23:20:40.373482515 +0530 IST m=+0.500738861
// Tick at 2022-11-26 23:20:40.873094027 +0530 IST m=+1.000350360
// Tick at 2022-11-26 23:20:41.373916893 +0530 IST m=+1.501173296
// Ticker Stopped
