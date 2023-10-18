package main

import "fmt"

func main() {
	fifo := FIFO{}
	cache := initCache(&fifo)

	cache.add("a", Key{value: "1", time: 1})
	cache.add("b", Key{value: "2", time: 2})
	// below entry will cause eviction (cap is 2)
	cache.add("c", Key{value: "3", time: 3})
	fmt.Println(cache)
	fmt.Println()

	lifo := LIFO{}
	cache.setEvictionAlgo(&lifo)

	// below entry will cause eviction
	cache.add("d", Key{value: "4", time: 4})
	fmt.Println(cache)
}

//Eviction by FIFO strategy
//&{map[b:{2 2} c:{3 3}] 0x5510b0 2 2}
//
//Eviction by LIFO strategy
//&{map[b:{2 2} d:{4 4}] 0x5510b0 2 2}
