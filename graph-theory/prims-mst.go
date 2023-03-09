package main

import (
	"fmt"
	"log"
)

type Node struct {
	Name interface{}
	Distance int
}

type Edge struct {
	from interface {}
	to interface{}
	distance int
}

type PriorityQueue struct {
	Edges []interface{}
}

func main() {
	edges := []Edge {
		Edge { from: 0, to: 1, distance: 10},
		Edge { from: 0, to: 2, distance: 1},
		Edge { from: 0, to: 3, distance: 4},
		Edge { from: 1, to: 2, distance: 3},
		Edge { from: 1, to: 4, distance: 0},
		Edge { from: 2, to: 5, distance: 8},
		Edge { from: 2, to: 3, distance: 2},
		Edge { from: 3, to: 5, distance: 2},
		Edge { from: 3, to: 6, distance: 7},
		Edge { from: 4, to: 5, distance: 1},
		Edge { from: 4, to: 7, distance: 8},
		Edge { from: 5, to: 7, distance: 9},
		Edge { from: 5, to: 6, distance: 6},
		Edge { from: 6, to: 7, distance: 12},
	}
	graph := buildGraph(edges)
	primsMST(graph)

}

func buildGraph(edges []Edge) map[interface{}][]interface{} {
	graph := make(map[interface{}][]interface{}) 
	for _, edge := range edges {
		if _, ok := graph[edge.from]; !ok {
			graph[edge.from] = []interface{}{}
		}
		if _, ok := graph[edge.to]; !ok {
			graph[edge.to] = []interface{}{}
		}
		graph[edge.from] = append(graph[edge.from], &Node { Name: edge.to, Distance: edge.distance})
		graph[edge.to] = append(graph[edge.to], &Node { Name: edge.from, Distance: edge.distance})
	}
	return graph
}

func primsMST(graph map[any][]any) {
	q := &PriorityQueue{}
	visited := make(map[interface{}]bool)
	start := 0
	var path []interface{} 

	// Initial 
	visited[0] = true
	for _, neighbor := range graph[start] {
		neighbor := cvNode(neighbor)
		q.Enqueue(&Edge { from: 0, to: neighbor.Name, distance: neighbor.Distance})
	}
	
	for len(q.Edges) > 0 {
		minEdge := cvEdge(q.Dequeue())

		if _, ok := visited[minEdge.to]; !ok {
			path = append(path, *minEdge)
		}
		visited[minEdge.to] = true

		for _, neighbor := range graph[minEdge.to] {
			neighbor := cvNode(neighbor)
			if _, ok := visited[neighbor.Name]; ok {
				continue
			}
			q.Enqueue(&Edge { from: minEdge.to, to: neighbor.Name, distance: neighbor.Distance})
		}

	}
	fmt.Println(path)
}

func (pq *PriorityQueue)Enqueue(e *Edge) {
	pq.Edges = append(pq.Edges, e)
	lastParentNode := len(pq.Edges) / 2 - 1
	buildMinHeap(pq.Edges, lastParentNode)
}

func (pq *PriorityQueue)Dequeue() any {
	if len(pq.Edges) > 0 {
		swap(pq.Edges, 0, len(pq.Edges) - 1)
		edge := pq.Edges[len(pq.Edges) - 1]
		pq.Edges = pq.Edges[:len(pq.Edges) - 1]
		rebuildHeapMin(pq.Edges, 0)
		return edge
	}
	return nil
}

func buildMinHeap(arr []interface{}, parentNode int) {
	for i := parentNode; i >= 0 ; i -- {
		leftNode := 2*i + 1
		rightNode := 2*i + 2
		largest := i
		n := len(arr)

		if leftNode < n && getValueCmp(arr[largest]) > getValueCmp(arr[leftNode]) {
			largest = leftNode
		}
	
		if rightNode < n && getValueCmp(arr[largest]) > getValueCmp(arr[rightNode]) {
			largest = rightNode
		}
	
		if largest != i {
			swap(arr, i, largest)
			rebuildHeapMin(arr, largest)
		}
	}
}

func rebuildHeapMin(arr []interface{}, parentNode int) {
	leftNode := 2*parentNode + 1
	rightNode := 2*parentNode + 2
	largest := parentNode
	n := len(arr)

	if leftNode < n && getValueCmp(arr[largest]) > getValueCmp(arr[leftNode]) {
		largest = leftNode
	}

	if rightNode < n && getValueCmp(arr[largest]) > getValueCmp(arr[rightNode]) {
		largest = rightNode
	}

	if largest != parentNode {
		swap(arr, parentNode, largest)
		rebuildHeapMin(arr, largest)
	}

}

func getValueCmp(i interface{}) int {
	switch v := i.(type) {
	case *Edge:
		return v.distance
	default:
		log.Panic("Type is invalid.")
		return 0
	}
}

func cvEdge(i interface{}) *Edge {
	switch v := i.(type) {
		case *Edge:
			return v
		default:
			log.Panic("Type is invalid.")
			return nil
	}
}

func cvNode(i interface{}) *Node {
	switch v := i.(type) {
		case *Node:
			return v
		default:
			log.Panic("Type is invalid.")
			return nil
	}
}

func swap(arr []interface{}, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
} 