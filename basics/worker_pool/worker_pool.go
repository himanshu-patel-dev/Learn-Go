package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	// for will wait on jobs channel until it get some data or it gets closed
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		// This starts up 3 workers, initially blocked because there are
		// no jobs yet.
		go worker(w, jobs, results)
	}
	// Here we send 5 jobs and then close that channel to indicate
	// that’s all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// Finally we collect all the results of the work. This also ensures
	// that the worker goroutines have finished. An alternative way to wait
	// for multiple goroutines is to use a WaitGroup.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

// $ time go run basics/worker_pool/worker_pool.go
// Worker 3 started job 1
// Worker 1 started job 2
// Worker 2 started job 3
// Worker 2 finished job 3
// Worker 2 started job 4
// Worker 3 finished job 1
// Worker 3 started job 5
// Worker 1 finished job 2
// Worker 3 finished job 5
// Worker 2 finished job 4

// real    0m2.209s
// user    0m0.221s
// sys     0m0.133s

// Our running program shows the 5 jobs being executed by various workers.
// The program only takes about 2 seconds despite doing about 5 seconds of
// total work because there are 3 workers operating concurrently.
