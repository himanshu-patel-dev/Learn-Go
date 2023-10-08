package main

import "fmt"

type DSU struct {
	data []int
}

func (dsu *DSU) init(n int) {
	for i := 0; i < n; i++ {
		dsu.data[i] = i
	}
}

func (dsu *DSU) find(x int) int {
	if dsu.data[x] != x {
		dsu.data[x] = dsu.find(dsu.data[x])
	}
	return dsu.data[x]
}

// union joins two different groups
// it returns true if two different group are attached
// false in case a link bw x and y already exists
func (dsu *DSU) union(x, y int) bool {
	rootx := dsu.find(x)
	rooty := dsu.find(y)
	if rootx == rooty {
		return false
	} else {
		dsu.data[rootx] = rooty
		return true
	}
}

// setLen return the unique groups in set
// if make a call to find() for each element
func (dsu *DSU) setLen() int {
	count := make(map[int]bool)
	for _, num := range dsu.data {
		x := dsu.find(num)
		count[x] = true
	}
	return len(count)
}

func main() {
	dsu := DSU{data: make([]int, 5)}
	dsu.init(5)

	dsu.union(0, 1)
	dsu.union(0, 2)
	dsu.union(3, 4)
	fmt.Println("--", dsu.setLen(), "--")

	fmt.Println(dsu.find(0))
	fmt.Println(dsu.find(1))
	fmt.Println(dsu.find(2))
	fmt.Println(dsu.find(4))
	fmt.Println(dsu.find(3))
}

// 2
// 2
// 2
// 4
// 4
