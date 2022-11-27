// Here we’ll look at using the sync/atomic package for atomic counters
// accessed by multiple goroutines.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// unsigned integer to represent our (always-positive) counter
	var ops uint64
	// a wait group to let main thread wait for all other goroutines
	var wg sync.WaitGroup

	// We’ll start 50 goroutines that each increment the counter
	// exactly 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// add 1 to count in atoming way
				// only one go routine will update at a time
				atomic.AddUint64(&ops, 1) // atomic operation
				// ops += 1 // non atomic operation
			}
			wg.Done() // every go routine when complete will notify waitgroup
		}()
	}

	// wait untile all goroutine are done
	wg.Wait()
	fmt.Println("Final Value : ", ops)
}

// with atomic operations: atomic.AddUint64(&ops, 1)
// Final Value :  50000

// without atomic operations
// Final Value :  33108
