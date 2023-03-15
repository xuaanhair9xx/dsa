package main

import(
	"fmt"
)

func main() {
	arr := []int {1, 3, 4, 0, 5, 0, 9, 2, 15, 1, 16}
	fmt.Println(longestPathSubproblem(arr));
	recursive(arr, len(arr))
	bottomUp(arr, len(arr))
	optimal(arr, len(arr))
}

func recursive(arr []int, n int) {	
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 1
		}
		if arr[i] > arr[i-1] {
			return 1 + dfs(i-1)
		} else {
			return dfs(i-1)	
		}
	}
	fmt.Println("Recursive approach: ", dfs(n-1))
}

func bottomUp(arr []int, n int) {
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i - 1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
	fmt.Println("Bottom up approach: ", dp[n-1])
}

func optimal(arr []int, n int) {
	v1 := 1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i - 1] {
			v1 ++
		}
	}
	fmt.Println("Optimal: ", v1)
}