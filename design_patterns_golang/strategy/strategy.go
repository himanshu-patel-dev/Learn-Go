package main

import (
	"fmt"
	"math"
)

type EvictionAlgo interface {
	evict(c *Cache)
}

// FIFO strategy
type FIFO struct {
}

func (f *FIFO) evict(c *Cache) {
	// write the logic to remove key from cache c
	// using FIFO
	minTime, minKey := math.MaxInt, ""
	for k, v := range c.storage {
		if v.time < minTime {
			minTime = v.time
			minKey = k
		}
	}
	delete(c.storage, minKey)

	fmt.Println("Eviction by FIFO strategy")
}

// LIFO strategy
type LIFO struct {
}

func (f *LIFO) evict(c *Cache) {
	// write the logic to remove key from cache c
	// using LIFO
	maxTime, maxKey := math.MinInt, ""
	for k, v := range c.storage {
		if v.time > maxTime {
			maxTime = v.time
			maxKey = k
		}
	}
	delete(c.storage, maxKey)

	fmt.Println("Eviction by LIFO strategy")
}
