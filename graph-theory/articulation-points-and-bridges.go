package main

import (
	"fmt"
	"time"
)

type Bridge struct {
	from interface {}
	to interface{}
}

var low map[interface{}]int
var id int = -1
var bridges []Bridge

func main() {
	low = make(map[interface{}]int)

	graph := map[interface{}][]interface{} { 
		"0": {"1"},
		"1": {"2"},
		"2": {"0", "5", "3"},
		"3": {"4"},
		"4": {},
		"5": {"6"},
		"6": {"7"},
		"7": {"8"},
		"8": {"5"},
	}
	arti(graph)
}



func arti(graph map[any][]any) {

	visited := make(map[any]bool)

	for at, _ := range graph {
		if _, ok := visited[at]; !ok {
			dsa(graph, at, -1, visited)
		}
	}

	time.Sleep(5 * time.Second)
	fmt.Println(low)
	fmt.Println(bridges)
}

func dsa(graph map[any][]any, at any, parent any, visited map[any]bool) {
	id = id + 1
	low[at] = id
	visited[at] = true

	for _, to := range graph[at] {
		if _, ok := visited[to]; !ok {
			dsa(graph, to, at, visited)
			low[at] = min(low[at], low[to])
			if low[at] < low[to] {
				bridges = append(bridges, Bridge {from: at, to: to})
			}
		} else {
			low[at] = min(low[at], low[to])
		}
	}
	return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}