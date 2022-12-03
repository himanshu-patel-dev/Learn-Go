package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Intn returns a random int n, 0 <= n < 100
	fmt.Println(rand.Intn(100), ",", rand.Intn(100))
	// 81 , 87	-- same value each time you exec the program

	// rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
	fmt.Println(rand.Float64())
	// 0.6645600532184904

	// can be used to generate random floats in other ranges,
	// for example 5.0 <= f' < 10.0
	fmt.Println(rand.Float64()*5 + 5)

	// The default number generator is deterministic, so it’ll produce
	// the same sequence of numbers each time by default. To produce
	// varying sequences, give it a seed that changes. Note that this
	// is not safe to use for random numbers you intend to be secret;
	// use crypto/rand for those.
	// But this is not thread safe
	source := rand.NewSource(time.Now().Unix())
	rand_generator := rand.New(source)

	fmt.Println(rand_generator.Intn(100))
	// 7 -- changes each time we exec the program

	// If you seed a source with the same number, it produces the same
	// sequence of random numbers.
	source1 := rand.NewSource(11)
	rand_gen1 := rand.New(source1)
	// two output below remain same as the seed value given to NewSource
	// reamin same
	fmt.Println(rand_gen1.Intn(100)) // 60
	fmt.Println(rand_gen1.Intn(100)) // 51

	// same seed means same sequence
	source2 := rand.NewSource(11)
	rand_gen2 := rand.New(source2)
	fmt.Println(rand_gen2.Intn(100)) // 60
	fmt.Println(rand_gen2.Intn(100)) // 51
}
