package main

import (
	"fmt"
	"errors"
	"log"
)

type Node struct {
	Name interface{}
	Distance int
}

func main() {
	dijkstra()
}

func dijkstra() {
	graph := map[interface{}][]interface{} { 
		0 : { Node { Name: 1 , Distance: 4 }, Node { Name: 2 , Distance: 1 }},
		1 : { Node { 3 , 1 } },
		2 : { Node { 1 , 2 }, Node { 3 , 5 }},
		3 : { Node { 4, 3 } },
	}
	var start interface{} = 0
	//var end interface{} = 4

	dist := make(map[interface{}]int)
	visited := make(map[interface{}]bool)
	fmt.Println(graph)

	queue := Queue { Size: 30 }

	startNode := Node { Name: start, Distance: 0 }
	queue.Enqueue(startNode)
	dist[cv(startNode).Name] = 0
	for queue.GetLength() > 0 {
		node := queue.Dequeue()
		visited[node] = true

		for _, neighbor := range graph[cv(node).Name] {
			if _, ok := visited[neighbor]; ok {
				continue
			} 
			newDist := dist[cv(node).Name] + cv(neighbor).Distance
			fmt.Printf("%v = %v + %v \n", newDist, dist[cv(node).Name], cv(neighbor).Distance)

			if _, ok := dist[cv(neighbor).Name]; !ok ||  dist[cv(neighbor).Name] > newDist {
			 	dist[cv(neighbor).Name] = newDist
			 	queue.Enqueue(neighbor)
		    } 
		}
	}

	fmt.Println("dist:", dist)

}

func cv(i interface {}) Node {
	switch v := i.(type) {
		case Node:
			return v
		default:
			log.Panic("Type is invalid.")
			return Node {}
	}
}

type Queue struct {
    Elements []interface {}
    Size     int
}

func (q *Queue) Enqueue(elem interface {}) {
    if q.GetLength() == q.Size && q.Size != 0 {
        fmt.Println("Overflow")
        return
    }
    q.Elements = append(q.Elements, elem)
}
  
func (q *Queue) Dequeue() interface {} {
    if q.IsEmpty() {
        fmt.Println("UnderFlow")
        return nil
    }
    element := q.Elements[0]
    if q.GetLength() == 1 {
        q.Elements = nil
        return element
    }
    q.Elements = q.Elements[1:]
    return element
}
  
func (q *Queue) GetLength() int {
    return len(q.Elements)
}
  
func (q *Queue) IsEmpty() bool {
    return len(q.Elements) == 0
}
  
func (q *Queue) Peek() ( interface {}, error) {
    if q.IsEmpty() {
        return nil, errors.New("empty queue")
    }
    return q.Elements[0], nil
}