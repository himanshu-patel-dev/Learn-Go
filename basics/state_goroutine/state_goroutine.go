// In the previous example we used explicit locking with mutexes to
// synchronize access to shared state across multiple goroutines.
// Another option is to use the built-in synchronization features of
// goroutines and channels to achieve the same result.

// This channel-based approach aligns with Go’s ideas of sharing memory
// by communicating and having each piece of data owned by exactly
// 1 goroutine.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// In this example our state will be owned by a single goroutine.
// This will guarantee that the data is never corrupted with concurrent
// access. In order to read or write that state, other goroutines will
// send messages to the owning goroutine and receive corresponding
// replies. These readOp and writeOp structs encapsulate those requests
// and a way for the owning goroutine to respond.

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// counters to keep count in atomic way
	var readOps uint64
	var writeOps uint64

	// channel to let go routines comm with each other
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// consumer go routine
	go func() {
		// kind of cache to keep record of writes
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// spin up 100 read go routines
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// read query input - a random int
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				// send this to read channel
				reads <- read
				// block until consumer return a response
				<-read.resp
				// inc the counter in atomic way
				atomic.AddUint64(&readOps, 1)
				// produce read query in every milli sec
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// spin 10 read go routines
	for w := 0; w < 10; w++ {
		go func() {
			for {
				// write query input - write a random int
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				// send this to write channel
				writes <- write
				// block until consumer return a response
				<-write.resp
				// inc the counter in atomic way
				atomic.AddUint64(&writeOps, 1)
				// producer write query in every milli sec
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

// readOps: 79005
// writeOps: 7960
