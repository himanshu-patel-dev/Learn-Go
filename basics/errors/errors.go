package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 0 {
		// create a new instance of error with given message
		return -1, errors.New("Can't use 0 as input")
	}
	// nil means no error
	return arg + 1, nil
}

// a custom type with custom error
type argError struct {
	arg  int
	prob string
}

// custom error defined, return a string explaining error
func (a *argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.prob)
}

// f2 use custom error type
func f2(arg int) (int, error) {
	if arg == 0 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 1, nil
}

func main() {
	for _, i := range []int{7, 0} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 0} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// cast the error in it's actual struct type
	// then access the fields
	_, e := f2(0)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

// f1 worked: 8
// f1 failed: Can't use 0 as input
// f2 worked: 8
// f2 failed: 0 - can't work with it
// 0
// can't work with it
