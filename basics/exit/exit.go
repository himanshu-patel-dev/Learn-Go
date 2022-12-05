// os.Exit to immediately exit with a given status
package main

import (
	"fmt"
	"os"
)

func main() {
	// defers will not be run when using os.Exit, so this
	// fmt.Println will never be called.
	defer fmt.Println("Exiting")
	// Exit with status 3
	os.Exit(3)
}

// usign go run
// exit status 3

// using build from go build
// $ ./exit
// $ echo $?
// 3
