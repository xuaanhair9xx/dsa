package main

import(
	"fmt"
	//"time"
)
var path []interface{}

func main() {

	graph := map[interface{}][]interface{} { 
		"0": {"1", "2", "3"},
		"1": {"0", "2"},
		"2": {"0", "1"},
		"4": {"3"},
		"3": {"0", "4"},
	}
	findEulerianPath(graph)
}

func findEulerianPath(graph map[any][]any) {
	in, out := countInOutDegree(graph);
	fmt.Println("in:", in)
	fmt.Println("out:", out)

	if graphHasEulerianPath(graph, in, out) {
		start := findStartNode(graph, in, out)
		fmt.Println("start", start)
		dfs(start, graph, in, out)
		fmt.Println("path:", path)

	}
}

func countInOutDegree(graph map[any][]any) (map[any]int, map[any]int) {
	out := make(map[interface{}]int)
	in := make(map[interface{}]int)

	for node, neighbors := range graph {
		for _, neighbor := range neighbors {
			if count, ok := out[node]; ok {
				out[node] = count + 1
			} else {
				out[node] = 1
			}

			if count, ok := in[neighbor]; ok {
				in[neighbor] = count + 1
			} else {
				in[neighbor] = 1
			}
		}
	}
	return in, out
}

func graphHasEulerianPath(graph map[any][]any, in, out map[any]int) bool {
	startNode := 0
	endNode := 0
	for i, _ := range graph   {
		if out[i] - in[i] > 1 || in[i] - out[i] > 1 {
			return false
		} else if out[i] - in[i] == 1 {
			startNode ++
		} else if in[i] - out[i] == 1 {
			endNode ++
		}
	}
	return (endNode == 0 && startNode == 0) || (endNode == 1 && startNode == 1)
}

func findStartNode(graph map[any][]any, in, out map[any]int) any {
	var start interface{} = "0"
	for i, _ := range graph {
		if out[i] - in[i] == 1 {
			return i
		}
		if out[i] > 0 {
			start = i
		}
	}
	return start
}

func dfs(at any, graph map[any][]any, in , out map[any]int) {
	for out[at] != 0 {
		out[at] = out[at] - 1
		next_edge := graph[at][out[at]]
		dfs(next_edge, graph, in , out)
	}
	path = append(path, at)
}