package main

import(
	"fmt"
)

func main() {
	arr := []int{9,1,1,9}
	n := len(arr)
	fmt.Println(dfs(n - 1, arr))
}
func dfs(i int, a []int) int {
	if (i < 0) {
		return 0
	}
	return max(a[i] + dfs(i - 2, a), dfs(i - 1, a))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}