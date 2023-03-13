package main

import (
	"fmt"
)

var arr []int
func main() {
	arr = []int {3, 34, 4, 12, 8, 2}
	sum := 30
	n := len(arr)
	recursive(n, sum)
	topDown(n, sum)
}

func topDown(n, sum int) {
	var dfs func(int, int) bool
	memo := make(map[int]bool)

	dfs = func(i int, sum int) bool {
		if _, ok := memo[i]; ok {
			return memo[i]
		}
		if sum == 0 {
			memo[i] = true
			return true
		}
		if i < 0 {
			return false
		}
		if arr[i] <= sum {
			memo[i] = dfs(i - 1, sum  - arr[i]) || dfs(i-1, sum)
		}
		memo[i] = dfs(i-1, sum)
		return memo[i]
	}
	fmt.Println("Top down approach: ", dfs(n-1, sum))
}

func recursive(n, sum int) {
	var dfs func(int, int) bool
	dfs = func(i int, sum int) bool {
		if sum == 0 {
			return true
		}
		if i < 0 {
			return false
		}
		if arr[i] <= sum {
			return dfs(i - 1, sum  - arr[i]) || dfs(i-1, sum)
		}
		return dfs(i-1, sum)
	}
	fmt.Println("Recursive approach: ", dfs(n-1, sum))
}