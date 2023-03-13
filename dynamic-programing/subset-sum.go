package main

import (
	"fmt"
)

var arr []int
func main() {
	arr = []int {3, 34, 4, 12, 8, 2}
	sum := 30
	n := len(arr) - 1
	fmt.Println(dfs(n, sum))
}

func dfs(i, sum int) bool {

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