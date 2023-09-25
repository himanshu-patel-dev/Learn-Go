package main

import (
	"fmt"
	"math"
)

func main() {
	// Abs - return absolute value of float
	fmt.Println("--- Abs ---")
	fmt.Println(math.Abs(-2))
	fmt.Println(math.Abs(2.5))
	// Abs(±Inf) = +Inf
	// Abs(NaN) = NaN

	// round off a float value to upper
	// integer value
	// Ceil(±0) = ±0
	// Ceil(±Inf) = ±Inf
	// Ceil(NaN) = NaN
	fmt.Println("--- Ceil ---")
	fmt.Println(math.Ceil(1.4)) // 2
	fmt.Println(math.Ceil(5))   // 5

	// round off a float value to lower
	// integer value
	// Ceil(±0) = ±0
	// Ceil(±Inf) = ±Inf
	// Ceil(NaN) = NaN
	fmt.Println("--- Floor ---")
	fmt.Println(math.Floor(1.4)) // 1
	fmt.Println(math.Floor(5))   // 5

	// represent +ve and -ve Inf
	fmt.Println("--- Inf ---")
	pos_inf := math.Inf(1)
	neg_inf := math.Inf(-1)
	fmt.Println(neg_inf) // -Inf
	fmt.Println(pos_inf) // +Inf

	// check if the value is Inf
	fmt.Println(math.IsInf(pos_inf, 1))  // true
	fmt.Println(math.IsInf(neg_inf, -1)) // true

	fmt.Println("--- NaN ---")
	// create and check for NaN - not a number
	// NaN (''not a number''): the result of mathematically
	// dubious operations as 0/0 or Sqrt(-1) .
	nan := math.NaN()
	nan1, nan2 := math.NaN(), math.NaN()
	fmt.Println(math.IsNaN(nan))                        // true
	fmt.Println(nan1 == nan2, nan1 < nan2, nan1 > nan2) // false false false

	fmt.Println("--- Log ---")
	// natural log
	// Log(+Inf) = +Inf
	// Log(0) = -Inf
	// Log(x < 0) = NaN
	// Log(NaN) = NaN

	fmt.Println(math.Log(1))      // 0
	fmt.Println(math.Log(2.7183)) // 1.0000066849139877

	// decimal log, base 10
	fmt.Println(math.Log10(10))  // 1
	fmt.Println(math.Log10(100)) // 2

	// log 2 base
	fmt.Println(math.Log2(256)) // 8

	fmt.Println("--- Max Min Mod ---")
	// Max(x, +Inf) = Max(+Inf, x) = +Inf
	// Max(x, NaN) = Max(NaN, x) = NaN
	// Max(+0, ±0) = Max(±0, +0) = +0
	// Max(-0, -0) = -0
	fmt.Println(math.Min(2.22, 3.33)) // 2.22
	fmt.Println(math.Max(2.22, 3.33)) // 3.33

	// Mod returns the floating-point remainder of x/y
	// Mod(±Inf, y) = NaN
	// Mod(NaN, y) = NaN
	// Mod(x, 0) = NaN
	// Mod(x, ±Inf) = x
	// Mod(x, NaN) = NaN
	fmt.Println(math.Mod(7, 4))   // 3
	fmt.Println(math.Mod(7.5, 4)) // 3.5

	fmt.Println("--- Pow ---")
	// Pow(x, ±0) = 1 for any x
	// Pow(1, y) = 1 for any y
	// Pow(x, 1) = x for any x
	// Pow(NaN, y) = NaN
	// Pow(x, NaN) = NaN
	// Pow(±0, y) = ±Inf for y an odd integer < 0
	// Pow(±0, -Inf) = +Inf
	// Pow(±0, +Inf) = +0
	// Pow(±0, y) = +Inf for finite y < 0 and not an odd integer
	// Pow(±0, y) = ±0 for y an odd integer > 0
	// Pow(±0, y) = +0 for finite y > 0 and not an odd integer
	// Pow(-1, ±Inf) = 1
	// Pow(x, +Inf) = +Inf for |x| > 1
	// Pow(x, -Inf) = +0 for |x| > 1
	// Pow(x, +Inf) = +0 for |x| < 1
	// Pow(x, -Inf) = +Inf for |x| < 1
	// Pow(+Inf, y) = +Inf for y > 0
	// Pow(+Inf, y) = +0 for y < 0
	// Pow(-Inf, y) = Pow(-0, -y)
	// Pow(x, y) = NaN for finite x < 0 and finite non-integer y
	fmt.Println(math.Pow(2, 3))  // 8
	fmt.Println(math.Pow(2, 0))  // 1
	fmt.Println(math.Pow(0, 2))  // 0
	fmt.Println(math.Pow(0, -1)) // +Inf
	fmt.Println(math.Pow(0, 0))  // 1 - this should be undefined

	// Pow10 returns 10**n, the base-10 exponential of n.
	// Pow10(n) =    0 for n < -323
	// Pow10(n) = +Inf for n > 308
	fmt.Println(math.Pow10(2)) // 100

	fmt.Println("--- Remainder ---")
	// Remainder(±Inf, y) = NaN
	// Remainder(NaN, y) = NaN
	// Remainder(x, 0) = NaN
	// Remainder(x, ±Inf) = x
	// Remainder(x, NaN) = NaN
	fmt.Println(math.Remainder(100, 30)) // 10

	fmt.Println("--- Round ---")
	// Round returns the nearest integer, rounding half away from zero.
	// RoundToEven(±0) = ±0
	// RoundToEven(±Inf) = ±Inf
	// RoundToEven(NaN) = NaN
	fmt.Println(math.Round(11.5)) // 12
	fmt.Println(math.Round(11.2)) // 11
}
