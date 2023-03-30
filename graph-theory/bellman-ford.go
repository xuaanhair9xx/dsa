package main

import "fmt"

type edgeT interface {
	int | string
}

type weightT interface {
	int | float64
} 

type Edge[T edgeT, W weightT] struct {
	src T 
	dest T 
	weight W
}

type Graph struct {
	V int // number of vertices
	E int // number of edges
	Edges []*Edge[int, int]
}

const INT_MAX = 10000

func initialGraph(v, e int) *Graph {
	edges := make([]*Edge[int, int], e)
	return &Graph{V: v, E: e, Edges: edges}
}

func bellmanFord(graph *Graph, src int) {
	v, e := graph.V, graph.E 
	dist := make([]int, v)

	// Step 1: Initialize distances from src to all other
    // vertices as INFINITE
	for i := 0; i < v; i ++ {
		dist[i] = INT_MAX
	}
	dist[src] = 0

	// Step 2: Relax all edges |V| - 1 times. A simple
    // shortest path from src to any other vertex can have
    // at-most |V| - 1 edges
	for i := 1; i <= v - 1; i++ {
		for j := 0; j < e; j++ {
			if dist[graph.Edges[j].src] != INT_MAX && dist[graph.Edges[j].dest] > graph.Edges[j].weight + dist[graph.Edges[j].src] {
				dist[graph.Edges[j].dest] = graph.Edges[j].weight + dist[graph.Edges[j].src]
			}
		}
	}

	// Step 3: check for negative-weight cycles.  The above
    // step guarantees shortest distances if graph doesn't
    // contain negative weight cycle.  If we get a shorter
    // path, then there is a cycle.
	for j := 0; j < e; j++ {
		if dist[graph.Edges[j].src] != INT_MAX && dist[graph.Edges[j].dest] > graph.Edges[j].weight + dist[graph.Edges[j].src] {
			fmt.Println("Graph contains negative weight cycle")
			return
		}
	}
	fmt.Println(dist)
}

func main() {
	graph := initialGraph(5, 8)

	// add edge 0-1 (or A-B in above figure)
	graph.Edges[0] = &Edge[int, int]{src: 0, dest:1, weight: -1}
  
    // add edge 0-2 (or A-C in above figure)
	graph.Edges[1] = &Edge[int, int]{src: 0, dest:2, weight: 4}
  
    // add edge 1-2 (or B-C in above figure)
	graph.Edges[2] = &Edge[int, int]{src: 1, dest:2, weight: 3}
  
    // add edge 1-3 (or B-D in above figure)
	graph.Edges[3] = &Edge[int, int]{src: 1, dest:3, weight: 2}
  
    // add edge 1-4 (or B-E in above figure)
	graph.Edges[4] = &Edge[int, int]{src: 1, dest:4, weight: 2}
  
    // add edge 3-2 (or D-C in above figure)
	graph.Edges[5] = &Edge[int, int]{src: 3, dest:2, weight: 5}
  
    // add edge 3-1 (or D-B in above figure)
	graph.Edges[6] = &Edge[int, int]{src: 3, dest:1, weight: 1}
  
    // add edge 4-3 (or E-D in above figure)
	graph.Edges[7] = &Edge[int, int]{src: 4, dest:3, weight: -3}

	bellmanFord(graph, 0)
}