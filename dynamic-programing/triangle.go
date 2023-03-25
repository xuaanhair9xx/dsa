package main

import "fmt"

func main() {
	triangle := [][]int {{2}, {3,4}, {6, 5, 7}, {4, 1, 8, 3}}
	way1(triangle)
}

func way1(triangle [][]int) {
	var dfs func(int, int) int
	n := len(triangle)
	dfs = func(i, j int) int {
		if i >= n {
			return 0
		}
		return min(triangle[i][j] + dfs(i+1,j), triangle[i][j] + dfs(i+1,j+1))
	}
	fmt.Println(dfs(0,0))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
