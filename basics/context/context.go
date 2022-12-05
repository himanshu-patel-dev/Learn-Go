// A Context carries deadlines, cancellation signals,
// and other request-scoped values across API boundaries
// and goroutines
package main

import (
	"fmt"
	"net/http"
	"time"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	// A context.Context is created for each request by the
	// net/http machinery, and is available with the
	// Context() method.
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Wait for a few seconds before sending a reply to the client.
	// This could simulate some work the server is doing. While
	// working, keep an eye on the context’s Done() channel for a
	// signal that we should cancel the work and return as soon as
	// possible.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// The context’s Err() method returns an error that explains
		// why the Done() channel was closed.
		err := ctx.Err()
		fmt.Println("server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/hello", HandleHello)
	http.ListenAndServe(":8090", nil)
}

// In first we wait for request to get completed while in second
// case we just close a waiting request before completion

// Client Console
// himanshu@workstation:~$ curl localhost:8090/hello
// hello
// himanshu@workstation:~$ curl localhost:8090/hello
// ^C

// Server Console
// server: hello handler started
// server: hello handler ended
// server: hello handler started
// server: context canceled
// server: hello handler ended
