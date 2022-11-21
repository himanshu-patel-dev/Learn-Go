package main

import "fmt"

// The main benefit that you get from using constants in
// your programs is the guarantee that their value will
// not change during program execution!

const s = "my-string"

func main() {
	fmt.Println(s)
	const n = 10
	if len(s) != 0 {
		fmt.Println(n) // 10
		const n = 20
		fmt.Println(n) // 20
	}
	//  n = 50 // can't reassign in same scope
	// const n = 50 // can't re declare in same scope
	fmt.Println(n) // 10
}
