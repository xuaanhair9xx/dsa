package main 

import(
	"fmt"
	"strconv"
	"log"
)

var order []interface {}

func main() {
	graph := map[interface{}][]interface{} { 
		"a": {"d"},
		"b": {"d"},
		"c": {"a", "b"},
		"d": {"g", "h"},
		"e": {"a", "d", "f"},
		"f": {"k", "j"},
		"g": {"i"},
		"h": {"i", "j"},
		"i": {"l"},
		"j": {"m", "l"},
	}
	topSort(graph)
}

func topSort(graph map[any][]any) {
	visited := make(map[any]bool)
	for node, _ := range graph {
		explore(graph, node, visited)
	}
	fmt.Println(order)
}

func explore(graph map[any][]any, node any, visited map[any]bool) {
	if _, ok := visited[node]; ok {
		return 
	}
	visited[node] = true
	for _, neighbor := range graph[node] {
		explore(graph, neighbor, visited)
	}
	order = append(order, node)
}

func isInListOrder(node any, order []string) bool {
	node = cvString(node)
	for _, n := range order {
		if n == node {
			return true
		}
	}
	return false
}

func cvString(i interface {}) string {
	switch v := i.(type) {
		case string:
			return v
		case int:
			vC := strconv.Itoa(v)
			return vC
		default:
			log.Panic("Type is invalid.")
			return "0"
	}
}