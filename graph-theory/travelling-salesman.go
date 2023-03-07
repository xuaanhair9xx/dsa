package main

import(
	"fmt"
)

func main() {
	graph := [][]int { { 0, 10, 15, 20 },{ 10, 0, 35, 25 },{ 15, 35, 0, 30 },{ 20, 25, 30, 0 } };
	start := 0;
	tsp(graph, start)
}

func tsp( graph [][]int, start int) {
	n := len(graph)
	memo := [n + 1][1 << (n + 1)]int{};
	setup(graph, memo, start, n)
}

func setup(graph [][]int, memo [10][10]int, start int, N int) {
	for i := 0; i < N; i ++ {
		if i == start {
			continue
		}
		fmt.Printf("%d \n", 1 << start | 1 << i)
		memo[i][1 << start | 1 << i] = graph[start][i]
	}
	fmt.Println(memo)
}