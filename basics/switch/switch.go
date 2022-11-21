package main

import (
	"fmt"
)

func main() {
	i := 24
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Don't Know")
	}
	//  Don't Know

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI("String")
	whatAmI(5)
	// I am bool
	// Don't know type string
	// I am int
}
