// In the previous example we looked at spawning external processes.
// We do this when we need an external process accessible to a running
// Go process. Sometimes we just want to completely replace the current
// Go process with another (perhaps non-Go) one. To do this we’ll use
// Go’s implementation of the classic exec function.

// Note that Go does not offer a classic Unix fork function. Usually
// this isn’t an issue though, since starting goroutines, spawning
// processes, and exec’ing processes covers most use cases for fork.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Go requires an absolute path to the binary we want to execute,
	// so we’ll use exec.LookPath to find it (probably /bin/ls).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires arguments in slice form (as opposed to
	// one big string). We’ll give ls a few common arguments.
	// Note that the first argument should be the program name.
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environment variables to use. Here
	// we just provide our current environment.
	env := os.Environ()

	// If this call is successful, the execution of our process
	// will end here and be replaced by the /bin/ls -a -l -h
	// process. If there is an error we’ll get a return value.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// When we run our program it is replaced by ls.

// $ go run exec_process/exec_process.go
// total 280K
// drwxrwxr-x 70 himanshu himanshu 4.0K Dec  5 09:46 .
// drwxrwxr-x  9 himanshu himanshu 4.0K Dec  2 08:39 ..
// drwxrwxr-x  2 himanshu himanshu 4.0K Nov 21 23:55 array
// drwxrwxr-x  2 himanshu himanshu 4.0K Nov 27 10:44 atomic_
