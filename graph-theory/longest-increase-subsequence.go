package main

import(
	"fmt"
	"errors"
)

func main() {
	arr := []int {1, 3, 4, 0, 5, 0, 9, 2}
	graph := buildGraph(arr)
	longestPath(graph)
}

func buildGraph(arr []int) map[int][]int{
	graph := make(map[int][]int)
	for i := 0; i < len(arr); i ++ {
		for j:= i + 1;j < len(arr);j++ {
			if (arr[i] < arr[j]) {
				graph[arr[i]] = append(graph[arr[i]], arr[j])
			}
		}
	} 
	return graph
}

func longestPath(graph map[int][]int) {
	max := 0
	fmt.Println(graph)
	for node, _ := range graph {
		// distance := exploreBfs(graph, node)
		distance := 0
		exploreDfs(graph, Node { Name: node, Distance: 1 }, &distance)
		if (max < distance) {
			max = distance
		}
	}
	fmt.Println(max)
}

func exploreDfs(graph map[int][]int, node Node, max *int) {
	for _, neighbor := range graph[node.Name] {
		exploreDfs(graph, Node { Name: neighbor, Distance: node.Distance + 1}, max)
	}
	if node.Distance > *max {
		*max = node.Distance
	}
}

func exploreBfs(graph map[int][]int, node int) int {
	queue := Queue { Size: 30 }
	queue.Enqueue(Node {Name: node, Distance: 1})
	var lastnode Node 
	for queue.GetLength() > 0 {
		lastnode = queue.Dequeue().(Node);
		for _, neighbor := range graph[lastnode.Name] {
			queue.Enqueue(Node { Name: neighbor, Distance: lastnode.Distance + 1})
		}
	}
	return lastnode.Distance
}

type Node struct {
	Name int
	Distance int
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

