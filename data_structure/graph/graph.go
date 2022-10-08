package main

import "fmt"

// graph structure
type Graph struct {
	vertices []*Vertex
}

// vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// add vertex
func (g *Graph) AddVertex(k int) {
	g.vertices = append(g.vertices, &Vertex{key: k})
}

// add edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		fmt.Printf("Edge already exists (%v->%v)\n", from, to)
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// get vertex
func (g *Graph) getVertex(key int) *Vertex {
	for _, v := range g.vertices {
		if v.key == key {
			return v
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}

// Print will print the adjacent list of each vertex of graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex: %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v, ", v.key)
		}
	}
	fmt.Println()
}

func main() {
	myGraph := &Graph{}
	// add 5 vertex
	for i := 0; i < 5; i++ {
		myGraph.AddVertex(i)
	}
	// add edges in vertices
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(4, 2)
	myGraph.AddEdge(3, 2)
	myGraph.AddEdge(7, 2) // invalid edge
	myGraph.AddEdge(4, 2) // existing edge
	myGraph.Print()
}
