package main

import (
	"fmt"
	"errors"
)

func permutationCount(graph map[string][]string, n int) {
	count := 0

	if (n == 1) {
		count = len(graph) 
	}

	for nodeName , _ := range graph {
		queue := QueueString { Size: 30 }
		newNode := Node { Name: nodeName, Distance: 0 }
		queue.Enqueue( newNode )

		CheckNeighbor: 
		for queue.GetLength() > 0 {
			xnode := queue.Dequeue()
			for _, neighborName := range graph[xnode.Name] {
				node := Node { 
					Name: neighborName, 
					Distance: xnode.Distance + 1 }
				queue.Enqueue(node)

				if node.Distance == n - 1 {
					count ++
				}

				if node.Distance == n {
					break CheckNeighbor
				}
			}
		}
	}
	fmt.Println(count)
}

func main() {
	graph := make(map[string][]string);
	graph = map[string][]string {
		"a" : { "e" },
		"e" : { "i", "a" },
		"i" : { "u", "o", "a", "e" },
		"o" : { "u", "i" },
		"u" : { "a" },
	} 
	permutationCount(graph, 1);
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