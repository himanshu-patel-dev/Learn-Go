package main

import "fmt"

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"

	fmt.Println(v1, v2, v3, v4)                    // add a space among each variable and a line after
	fmt.Print(v1, v2, v3, v4)                      // no space and no line in the end
	fmt.Println()                                  // just to add a line after the previous print end
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n") // in order to make in similar to other print we need to add literals manually
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)     // in order to get even fine control
}
