package main

import(
	"log"
	"fmt"
)

type Vertex struct {
	name interface{}
	distance int
}

type Edge struct {
	from interface {}
	to interface{}
	distance int
}

type Object struct {
	edge *Edge
	index interface{}
}

type IPQ struct {
	edges []interface{}
}

func main() {
	edges := []Edge {
		Edge { from: 0, to: 1, distance: 9},
		Edge { from: 0, to: 2, distance: 0},
		Edge { from: 0, to: 3, distance: 5},
		Edge { from: 0, to: 5, distance: 7},
		Edge { from: 1, to: 3, distance: -2},
		Edge { from: 1, to: 6, distance: 4},
		Edge { from: 1, to: 4, distance: 3},
		Edge { from: 2, to: 5, distance: 6},
		Edge { from: 3, to: 5, distance: 2},
		Edge { from: 3, to: 6, distance: 3},
		Edge { from: 4, to: 6, distance: 6},
		Edge { from: 5, to: 6, distance: 1},
	}
	graph := buildGraph(edges)
	edgerPrimsMST(graph)
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
		graph[edge.from] = append(graph[edge.from], &Vertex { name: edge.to, distance: edge.distance})
		graph[edge.to] = append(graph[edge.to], &Vertex { name: edge.from, distance: edge.distance})
	}
	return graph
}

func edgerPrimsMST(graph map[any][]any) {
	visited := make(map[interface{}]bool)
	ipq := &IPQ{}
	relaxEdgesAtNode(0, graph, visited, ipq)
	var path []interface {} 
	for ipq.GetLength() > 0 {
		current := ipq.Dequeue()
		if _, ok := visited[cvObject(current).index]; !ok {
			path = append(path, current)
		}
		relaxEdgesAtNode(cvObject(current).index, graph, visited, ipq)
	}
	for _, x := range path {
		fmt.Println(cvObject(x).edge)
	}
}

func relaxEdgesAtNode(currentIndex any, graph map[any][]any, visited map[any]bool, q *IPQ) {
	visited[currentIndex] = true
	for _, neighbor := range graph[currentIndex] {
		neighbor := cvVertex(neighbor)
		if _, ok := visited[neighbor.name]; ok {
			continue
		}
		object := &Object{
			edge: &Edge { from: currentIndex, to: neighbor.name, distance: neighbor.distance }, 
			index: neighbor.name,
		}
		i := q.IsContain(object)
		q.Enqueue(object, i)
	}
}

func (pq *IPQ)Enqueue(e any, index int) {
	if (index != -1) {
		if getValueCmp(pq.edges[index]) > getValueCmp(e) {
			pq.edges[index] = e
		} else {
			return
		}
	} else {
		pq.edges = append(pq.edges, e)
	}
	lastParentNode := len(pq.edges) / 2 - 1
	buildMinHeap(pq.edges, lastParentNode)
}

func (pq *IPQ)Dequeue() any {
	if len(pq.edges) > 0 {
		swap(pq.edges, 0, len(pq.edges) - 1)
		object := pq.edges[len(pq.edges) - 1]
		pq.edges = pq.edges[:len(pq.edges) - 1]
		rebuildHeapMin(pq.edges, 0)
		return object
	}
	return nil
}

func (pq *IPQ)GetLength() int {
	return len(pq.edges)
}

func (pq *IPQ)IsContain(e any) int {
	// Check edge is existed.
	index := -1
	for i, obj := range pq.edges {
		if cvObject(obj).index == cvObject(e).index {
			index = i
		}
	}
	return index
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
	case *Object:
		return v.edge.distance
	default:
		log.Panic("Type is invalid.")
		return 0
	}
}

func cvObject(i interface{}) *Object {
	switch v := i.(type) {
		case *Object:
			return v
		default:
			log.Panic("Type is invalid.")
			return nil
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

func cvVertex(i interface{}) *Vertex {
	switch v := i.(type) {
		case *Vertex:
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