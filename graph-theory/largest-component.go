package main

import (
	"fmt"
	"strconv"
)

func main() {
	var graph = make(map[int][]string)
	
	graph = map[int][]string {
		0 : {"8", "1", "5"},
		1 : {"0"},
		5 : {"0","8"},
		8 : {"0","5"},
		2 : {"3","4"},
		3 : {"2","4"},
	}
	largestComponent(graph)
}

func largestComponent(graph map[int][]string) {
	var visited = make(map[int]bool)
	largest := 0

	for node, _ := range graph {
		if _, ok := visited[node]; ok {
			continue
		}
		size := exploreSize(node, graph, visited)
		if size > largest {
			largest = size;
		}
	}
	fmt.Println(largest);
}

func exploreSize(current int, graph map[int][]string, visited map[int]bool) int {
	if _, ok := visited[current]; ok {
		return 0
	}
	visited[current] = true
	count := 1
	for _, node := range graph[current] {
		iNode, err := strconv.Atoi(node)
		if err != nil {
			panic(err)
		}
		ret := exploreSize(iNode, graph, visited)
		fmt.Printf("count: %d + ret: %d \n", count, ret)
		count = count + ret
	}
	fmt.Printf("end.....: %d \n", count)
	return count
}