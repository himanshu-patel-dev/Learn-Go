package main

import "fmt"

func main() {
	f := fmt.Println

	binary := 0b1111
	f(binary) // 15
	decimal := 15
	f(decimal)   // 15
	octal := 017 // = 1*8 + 7 = 15
	f(octal)     // 15
	f(1_234_567) // 1234567
	f(1_2_3_4)   // 1234

	// Rune literal represent characters and are surrounded by literals
	f('\141')       // 8-bit octal numbers
	f('\x61')       // 8-bit hexadecimal numbers
	f('\u0061')     // 16-bit hexadecimal numbers
	f('\U00000061') // 32-bit Unicode numbers
	f(string(97))   // a
	// all are 97, print string(97) = a

	f("Greetings and\n\"Salutations\"")
	//Greetings and
	//"Salutations"

	rawString := `Greetings and
	"Salutations"` // there is a tab preceding the second string, hence the print
	f(rawString)
	//Greetings and
	//	"Salutations"

	// type name - value range
	//int8			–128 to 127
	//int16			–32768 to 32767
	//int32			–2147483648 to 2147483647
	//int64			–9223372036854775808 to 9223372036854775807
	//uint8			0 to 255
	//uint16		0 to 65536
	//uint32		0 to 4294967295
	//uint64		0 to 18446744073709551615

	// On a 32-bit CPU, int is a 32-bit signed integer like
	// an int32. On most 64-bit CPUs, int is a 64-bit signed integer, just like an int64.
	// Because int isn’t consistent from platform to platform, it is a compile-time error to
	// assign, compare, or perform mathematical operations between an int and an int32
	// or int64 without a type conversion
}
