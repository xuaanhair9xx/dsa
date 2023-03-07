package main

import (
	"fmt"
	"errors"
)

func convertEdgesToGraph(edges [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}

func shortestPath(edges [][]string, start string, end string) {
	graph := convertEdgesToGraph(edges)
	queue := QueueString {Size: 30}
	var listDesNode []Node
	//minPath := len(graph)

	queue.Enqueue( Node { Name: start, Distance: 0 })
	visited := make(map[string]bool)

	for queue.GetLength() > 0 {
		xnode := queue.Dequeue()
		visited[xnode.Name] = true

		for _, nodeName := range (graph[xnode.Name]) {
			node := Node {Name: nodeName, Distance: xnode.Distance + 1}
			if _, ok := visited[nodeName]; ok {
				continue
			}
			queue.Enqueue(node)
			if (nodeName == end) {
				listDesNode = append(listDesNode, node)
			}
		}
	}

	fmt.Println(listDesNode);
}

func main() {
	edges := [][]string {
		{"w", "x"},
		{"x", "y"},
		{"z", "y"},
		{"z", "v"},
		{"w", "v"},
	}
	shortestPath(edges, "w", "z")
}

type Node struct {
	Distance int
	Name string
}

type QueueString struct {
    Elements []Node
    Size     int
}
func (q *QueueString) Enqueue(elem Node) {
    if q.GetLength() == q.Size {
        fmt.Println("Overflow")
        return
    }
    q.Elements = append(q.Elements, elem)
}
  
func (q *QueueString) Dequeue() Node {
    if q.IsEmpty() {
        fmt.Println("UnderFlow")
        return Node{}
    }
    element := q.Elements[0]
    if q.GetLength() == 1 {
        q.Elements = nil
        return element
    }
    q.Elements = q.Elements[1:]
    return element // Slice off the element once it is dequeued.
}
  
func (q *QueueString) GetLength() int {
    return len(q.Elements)
}
  
func (q *QueueString) IsEmpty() bool {
    return len(q.Elements) == 0
}
  
func (q *QueueString) Peek() (Node, error) {
    if q.IsEmpty() {
        return Node {}, errors.New("empty queue")
    }
    return q.Elements[0], nil
}
