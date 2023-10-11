package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Dijkstra struct {
	graph    [][]int
	distance []int
	vertexes int
}

// init takes data as the adj matrix of size v
// and v as count of vertex in graph
func (d *Dijkstra) init(data [][]int, v int) {
	d.graph = data
	d.vertexes = v
	d.distance = make([]int, v) // all distance are 0
}

// computeShortestPath compute shortest path from src
// vertex to all other vertex and
func (d *Dijkstra) computeShortestPath(src int) {
	/*
		Time Complexity = O()
	*/

	// all vertex are inf dist initially
	for i := range d.distance {
		d.distance[i] = math.MaxInt
	}
	// mark src vertex at dist 0
	d.distance[src] = 0
	// all vertices are unvisited - false by default
	var visited []bool = make([]bool, d.vertexes)
	var minHeap *MinHeap = &MinHeap{}
	heap.Init(minHeap)
	// push src vertex on min heap
	heap.Push(minHeap, &MinNode{distance: 0, node: src})

	// max iteration = no of edges = E = V*V (max)
	for minHeap.Len() > 0 {
		var minDistNode *MinNode = heap.Pop(minHeap).(*MinNode)

		// it this node is already visited (might be because
		// of lower distance from some other prevous processed
		// node)
		if visited[minDistNode.node] {
			continue
		}
		visited[minDistNode.node] = true
		d.distance[minDistNode.node] = minDistNode.distance

		// visit all unvisited neighbours
		for nbrNode, nbrDist := range d.graph[minDistNode.node] {
			// skip the vertices which are already visited, this prevent
			// next node to stop processing previous/parent node
			// same thing if there is no edge
			if visited[nbrNode] || nbrDist <= 0 {
				continue
			}
			var newDist int = nbrDist + minDistNode.distance
			// if nbrNode can be reached with lesser distance then
			// push this edge to heap to be visited in later iterations
			if newDist < d.distance[nbrNode] {
				heap.Push(minHeap, &MinNode{distance: newDist, node: nbrNode})
			}
		}
	}
}

func (d *Dijkstra) printDistance() {
	for i, dist := range d.distance {
		fmt.Printf(
			"Vertex = %d \t- Distance from source = %d\n", i, dist)
	}
}

func main() {
	var data [][]int = [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}
	var d *Dijkstra = &Dijkstra{}
	d.init(data, len(data))
	var sourceVertex int = 0
	d.computeShortestPath(sourceVertex)
	d.printDistance()
}

// -----------------------------------------------

type MinNode struct {
	distance int
	node     int
}

type MinHeap []*MinNode

func (m *MinHeap) Len() int { return len(*m) }
func (m *MinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}
func (m *MinHeap) Less(i, j int) bool {
	return (*m)[i].distance < (*m)[j].distance
}
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(*MinNode))
}
func (m *MinHeap) Pop() any {
	n, old := m.Len(), *m
	lastEle := old[n-1]
	*m = old[0 : n-1]
	return lastEle
}
