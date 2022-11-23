// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel’s receivers.

package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// worker go routine which consume the jobs sent on jobs channel and inform on done channel
	go func() {
		for {
			// the more value will be false if jobs has been closed and all values in the
			// channel have already been received
			job, more := <-jobs
			if more {
				fmt.Println("Received Jobs: ", job)
			} else {
				fmt.Println("Received All Jobs")
				done <- true
				return
			}
		}
	}()

	// send 3 jobs to channel and worker listening on this routine will consume them
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// close the channel after 3 jobs to signal the completion of work sent to worker
	close(jobs)
	fmt.Println("sent all jobs")

	<-done // wait for worker to finish and signal
}
