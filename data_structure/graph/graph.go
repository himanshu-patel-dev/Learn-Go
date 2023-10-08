package main

import (
	"fmt"
	"slices"
)

// Graph structure
type Graph struct {
	vertices int
	adjList  map[int][]int
}

// NewGraph return a pointer to a new graph
func NewGraph(n int) *Graph {
	return &Graph{
		vertices: n,
		adjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(src, dest int) {
	// skip invalid edge addition, comment below logic in case
	// we can have vertex number that don't depends on
	if g.vertices <= src || g.vertices <= dest {
		return
	}
	if !slices.Contains(g.adjList[src], dest) {
		g.adjList[src] = append(g.adjList[src], dest)
	}
	if !slices.Contains(g.adjList[dest], src) {
		g.adjList[dest] = append(g.adjList[dest], src)
	}
}

func (g *Graph) DFSUtil(src int, dest int, visited map[int]bool) bool {
	if visited[src] {
		return false
	}
	visited[src] = true
	// if src node is dest node, then we reached
	// the desired node return true
	if src == dest {
		return true
	}

	// iterate for all neighbour and dfs for dest node
	for _, nbr := range g.adjList[src] {
		if g.DFSUtil(nbr, dest, visited) {
			return true
		}
	}
	return false
}

func (g *Graph) DFS(src, dest int) {
	visited := make(map[int]bool)
	g.DFSUtil(src, dest, visited)
}

func (g *Graph) Print() {
	for node, nbr := range g.adjList {
		fmt.Println(node, nbr)
	}
}

func main() {
	myGraph := NewGraph(5)
	// add edges in vertices
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(4, 2)
	myGraph.AddEdge(3, 2)
	myGraph.AddEdge(7, 2) // invalid edge
	myGraph.AddEdge(4, 2) // existing edge

	myGraph.Print()
}

// 1 [2]
// 2 [1 4 3]
// 4 [2]
// 3 [2]
