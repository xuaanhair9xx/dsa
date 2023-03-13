package main

import(
	"fmt"
)

func main() {
	arr := []int{9,1,1,9}
	n := len(arr)
	recursive(n, arr)
	topDown(n, arr)
	bottomUp(n, arr)
}

// Recursive approach.
func recursive(n int, a []int) {
	var dfs func(int) int
	dfs = func(i int) int {
		if (i < 0) {
			return 0
		}
		return max(a[i] + dfs(i - 2), dfs(i - 1))
	}
	fmt.Println("Recursive approach: ", dfs(n-1))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Top down approach.
func topDown(n int, a []int) {
	memo := make(map[int]int)
	var dfs func(int) int
	dfs = func(i int) int {
		if _, ok := memo[i]; ok {
			return memo[i]
		}
		if (i < 0) {
			return 0
		}
		memo[i] = max(a[i] + dfs(i - 2), dfs(i - 1))
		return memo[i]
	}
	fmt.Println("Top down approach: ", dfs(n-1))
}

func bottomUp(n int, a []int) {
	var dp []int
	var v1, v2 int
	for i := 0; i < n; i++ {
		if i - 2 < 0 {
			v1 = 0
		} else {
			v1 = dp[i-2]
		}
		if i - 1 < 0 {
			v2 = 0
		} else {
			v2 = dp[i-1]
		}
		dp = append(dp, max(a[i] + v1, v2))
	}   
	fmt.Println("Bottom up approach: ", dp[len(dp)-1])
}