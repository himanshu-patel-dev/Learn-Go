package main

import (
	"fmt"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

// reading all the HTTP request headers and echoing
// them into the response body
func HandleHeaders(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// It sets up the default router in the net/http
	// package and takes a function as an argument
	http.HandleFunc("/hello", HandleHello)
	http.HandleFunc("/headers", HandleHeaders)

	http.ListenAndServe(":8080", nil)
}

// himanshu@workstation:~$ curl localhost:8080/hello
// Hello
// himanshu@workstation:~$ curl localhost:8080/headers
// User-Agent: curl/7.81.0
// Accept: */*
