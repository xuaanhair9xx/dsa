package main

import (
	"fmt"
	"errors"
)
// Connected component count

type Queue struct {
    Elements []int
    Size     int
}
func (q *Queue) Enqueue(elem int) {
    if q.GetLength() == q.Size {
        fmt.Println("Overflow")
        return
    }
    q.Elements = append(q.Elements, elem)
}
  
func (q *Queue) Dequeue() int {
    if q.IsEmpty() {
        fmt.Println("UnderFlow")
        return 0
    }
    element := q.Elements[0]
    if q.GetLength() == 1 {
        q.Elements = nil
        return element
    }
    q.Elements = q.Elements[1:]
    return element // Slice off the element once it is dequeued.
}
  
func (q *Queue) GetLength() int {
    return len(q.Elements)
}
  
func (q *Queue) IsEmpty() bool {
    return len(q.Elements) == 0
}
  
func (q *Queue) Peek() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("empty queue")
    }
    return q.Elements[0], nil
}


func ConnectedComponentCount(graph [][]int) {
	travesed := make(map[int]int)
	count := 0
	// new linked list
	queue := Queue {Size: 30}

	for i, _ := range graph {
		if _, ok := travesed[i]; ok {
			continue
		}

		travesed[i] = 1
		queue.Enqueue(i)
		for queue.GetLength() > 0 {
			check := queue.Dequeue()
			for _, neighbor := range graph[check] {
				if _, ok := travesed[neighbor]; ok {
					continue
				}
				queue.Enqueue(neighbor)
				travesed[neighbor] = 1
			}
		}
		count ++

	}

	fmt.Println(count)
}

func main()  {
	var graph = [][]int {
		{6,1,5},
		{0},
		{3, 4},
		{2, 4},
		{3, 2},
		{0, 6},
		{0, 5},
	}
	ConnectedComponentCount2(graph)
}

func ConnectedComponentCount2(graph [][]int) {
	visited := make(map[int]int)
	count := 0
	for node, _ := range graph {

		if _, ok := visited[node]; ok {
			continue
		}

		if explore(node, graph, visited) == true {
			count ++
		}
	}
	fmt.Println(count)
}

func explore(current int, graph [][]int, visited map[int]int) bool {
	if _, ok := visited[current]; ok {
		return false
	}
	visited[current] = 1
	for _, neighbor := range graph[current] {
		if (explore(neighbor, graph, visited) == false) {
			break 
		} 
	}
	return true
}