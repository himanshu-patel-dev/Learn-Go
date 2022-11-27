// we saw how to manage simple counter state using atomic operations.
// For more complex state we can use a mutex to safely access data
// across multiple goroutines.

package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; since we want to update it
// concurrently from multiple goroutines, we add a Mutex to synchronize
// access. Note that mutexes must not be copied, so if this struct is
// passed around, it should be done by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// this lock will lock the entire Container struct
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	// make a wait group to let main thread wait for go routines
	var wg sync.WaitGroup

	doIncrease := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		// inform wait group once done
		// this logic should be inbuild in func of goroutine
		wg.Done()
	}

	// declare to waitgroup to wait for 3 goroutine
	wg.Add(3)
	go doIncrease("a", 10000)
	go doIncrease("a", 10000)
	go doIncrease("b", 10000)

	wg.Wait()
	fmt.Println(c.counters) // map[a:20000 b:10000]
}
