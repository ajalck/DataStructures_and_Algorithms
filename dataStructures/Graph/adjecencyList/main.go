package main

import (
	"fmt"
	"sync"
)

//Graph represents an adjacency list graph

type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key      int
	adjacent []int
}

//AddVertex adds a vertex to the graph

func (g *Graph) AddVertex(k int) {
	if g.Contains(k) {
		fmt.Printf("Vertex %d already exists !", k)
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

//Contains checks the vertex's extistancy

func (g *Graph) Contains(k int) bool {
	for _, v := range g.vertices {
		if v.key == k {
			return true
		}
	}
	return false
}

//PrintGraph will print the adjecency list

func (g *Graph) PrintGraph() {
	fmt.Println("Adjecency list")
	for _, v := range g.vertices {
		fmt.Println(v.key, " --> ", v.adjacent)
	}
	fmt.Println()
}
func main() {
	test := &Graph{}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 10; i < 20; i++ {
			test.AddVertex(i)
		}
		wg.Done()
	}()
	wg.Wait()
	test.PrintGraph()
	test.AddVertex(15)


}
