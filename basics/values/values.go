package main

import "fmt"

func main() {
	fmt.Println("Hello" + "World")   // HelloWorld
	fmt.Println("2+2.5 =", 2+2.5)    // 2+2.5 = 4.5
	fmt.Println("5.0/3.0=", 5.0/3.0) // 5.0/3.0= 1.6666666666666667
	fmt.Println(true && false)       // false
	fmt.Println(true || false)       // true
	fmt.Println(!false)              // true
}
