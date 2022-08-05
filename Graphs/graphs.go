package main

import (
	"fmt"
	"log"
)

// Graph : represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex : represents a graph vertex
type Vertex struct {
	key int
	adjacent []*Vertex
}

// AddVertex : adds a Vertex to the Graph
func (g *Graph) AddVertex(k int) {
	var err error
	if !contains(g.vertices, k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
		return
	} else {
		err = fmt.Errorf("vertex %v not added it already exist", k)
	}

	if err != nil {
		log.Print(err.Error())
	}
}

// AddEdge : adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// check errors
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v)", from, to)
		log.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v --> %v)", from, to)
		log.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

// getVertex : returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains : returns true if it contains a duplicated vertex
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print : will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, va := range v.adjacent {
			fmt.Printf(" %v ", va.key)
		}
	}
	fmt.Println()
}

func main() {
	myGraph := &Graph{}

	for i := 0; i < 5; i++ {
		myGraph.AddVertex(i)
	}
	myGraph.AddVertex(0)
	
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(6, 2)

	myGraph.Print()
}