package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a wait group.

// This is the function we’ll run in every goroutine.
func worker(id int) {
	fmt.Println("Worker starting: ", id)
	time.Sleep(time.Second)
	fmt.Println("Worker done: ", id)
}

func main() {
	// This WaitGroup is used to wait for all the goroutines launched here to
	// finish. Note: if a WaitGroup is explicitly passed into functions,
	// it should be done by pointer.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Launch several goroutines and increment the WaitGroup counter for each.
		wg.Add(1)
		// Avoid re-use of the same i value in each goroutine closure.
		i := i

		// Wrap the worker call in a closure that makes sure to tell the
		// WaitGroup that this worker is done.
		// This way the worker itself does not have to be aware of the
		// concurrency primitives involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	// Block until the WaitGroup counter goes back to 0; all the
	// workers notified they’re done.
	wg.Wait()
}

// Note that this approach has no straightforward way to propagate errors from
// workers. For more advanced use cases, consider using the errgroup package.

// Worker starting:  5
// Worker starting:  1
// Worker starting:  4
// Worker starting:  2
// Worker starting:  3
// Worker done:  3
// Worker done:  2
// Worker done:  5
// Worker done:  4
// Worker done:  1

// real    0m1.228s
