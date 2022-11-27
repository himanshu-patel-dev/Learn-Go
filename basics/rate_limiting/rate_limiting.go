// Rate limiting is an important mechanism for controlling resource
// utilization and maintaining quality of service. Go elegantly supports
// rate limiting with goroutines, channels, and tickers.

package main

import (
	"fmt"
	"time"
)

func main() {
	// chan to receive request
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	// close channel to convey the finite number of request
	// and the reader don't keep waiting for new request
	close(request)

	// This limiter channel will receive a value every 200 milliseconds.
	// This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	for req := range request {
		<-limiter
		fmt.Println("Request per 200ms : ", req, time.Now())
	}
	// Request per 200ms :  1 2022-11-27 10:18:12.557212094
	// Request per 200ms :  2 2022-11-27 10:18:12.757612237
	// Request per 200ms :  3 2022-11-27 10:18:12.95689864
	// Request per 200ms :  4 2022-11-27 10:18:13.157260288
	// Request per 200ms :  5 2022-11-27 10:18:13.357621687

	// a channel with 3 as buffer limit
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	fmt.Println("---------------------------------------------------")

	// a go routine which have a time ticker
	// each time it ticks it allow limiter to accept one request
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// burstRecords depicts 5 instant request
	burstyRecords := make(chan int, 5)
	// fill the request channel with 5 request
	for i := 0; i < 5; i++ {
		burstyRecords <- i
	}
	close(burstyRecords)

	// we serve the first 3 immediately because of the burstable rate
	// limiting, then serve the remaining 2 with ~200ms delays each.
	for req := range burstyRecords {
		<-burstyLimiter
		fmt.Println("Request with burst: ", req, time.Now())
	}
	// Request with burst:  0 2022-11-27 10:40:30.481412411
	// Request with burst:  1 2022-11-27 10:40:30.481450889
	// Request with burst:  2 2022-11-27 10:40:30.481475443
	// Request with burst:  3 2022-11-27 10:40:30.681798332
	// Request with burst:  4 2022-11-27 10:40:30.882394142
}
