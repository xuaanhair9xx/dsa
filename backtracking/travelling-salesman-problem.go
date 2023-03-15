package main

import(
	"fmt"
)

const MAX_INT = 1000000
var visited []bool
var graph [][]int
var minCost int
var n int 

func main() {
	graph = [][]int { { 0, 10, 15, 20 },{ 10, 0, 35, 25 },{ 15, 35, 0, 30 },{ 20, 25, 30, 0 } };
	visited = make([]bool, len(graph))
	start := 0;
	visited[start] = true
	minCost = MAX_INT
	fmt.Println(tps(start, 0, start))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func tps(curpost int, curCost int, start int) int {
	if visitedAll() && graph[curpost][start] != 0 {
		minCost = min(minCost, curCost + graph[curpost][start])
		return minCost
	}
	for pos, cost := range graph[curpost] {
		if !visited[pos] {
			visited[pos] = true
			tps(pos, curCost + cost, start)
			visited[pos] = false
		}
	}
	return minCost
}

func visitedAll() bool {
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}