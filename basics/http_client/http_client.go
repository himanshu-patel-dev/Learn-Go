package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// it uses the http.DefaultClient object which
	// has useful default settings
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	// Response Status:  200 OK

	// Print the first 5 lines of the response body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; i < 5 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
}

// <!DOCTYPE html>
// <html>
//   <head>
//     <meta charset="utf-8">
//     <title>Go by Example</title>
