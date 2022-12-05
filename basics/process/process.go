package main

import (
	"fmt"
	"io"
	"os/exec"
)

// Sometimes our Go programs need to spawn other, non-Go processes.

func main() {
	// The exec.Command helper creates an object to
	// represent this external process.
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(">date")
	fmt.Println(string(dateOut))
	// >date
	// Mon Dec  5 09:27:56 AM IST 2022

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("Failed Executing ", err)
		case *exec.ExitError:
			fmt.Println("Command Exit ", e.ExitCode())
		default:
			panic(err)
		}
	}
	// Command Exit  1

	// Next we’ll look at a slightly more involved case where
	// we pipe data to the external process on its stdin
	// and collect the results from its stdout
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	// Here we explicitly grab input/output pipes, start the
	// process, write some input to it, read the resulting output,
	// and finally wait for the process to exit.
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
	// > grep hello
	// hello grep

	// Note that when spawning commands we need to provide an
	// explicitly delineated command and argument array, vs. being
	// able to just pass in one command-line string. If you want
	// to spawn a full command with a string, you can use
	// bash’s -c option
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -alh")
	fmt.Println(string(lsOut))
	// > ls -alh
	// total 52K
	// drwxrwxr-x  9 himanshu himanshu 4.0K Dec  2 08:39 .
	// drwxrwxr-x  9 himanshu himanshu 4.0K Nov 26 01:03 ..
	// drwxrwxr-x 70 himanshu himanshu 4.0K Dec  5 09:46 basics
}
