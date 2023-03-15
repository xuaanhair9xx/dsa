package main

import (
	"fmt"
)

var arr []int
func main() {
	arr = []int {3, 34, 4, 12, 8, 2}
	sum := 29
	n := len(arr)
	recursive(n, sum)
	topDown(n, sum)
	bottomUp(n, sum)
}

func topDown(n, sum int) {
	var dfs func(int, int) bool

	var memo = make([][]interface{}, n)
	for i := range memo {
		memo[i] = make([]interface{}, sum + 1)
	}

	dfs = func(i int, sum int) bool {
		if sum == 0 {
			return true
		}
		if i < 0 {
			return false
		}
		if memo[i][sum] != nil {
			return memo[i][sum].(bool)
		}
		if arr[i] <= sum {
			memo[i][sum] = dfs(i - 1, sum  - arr[i]) || dfs(i-1, sum)
		} else {
			memo[i][sum] = dfs(i-1, sum)
		}
		return memo[i][sum].(bool)
	}
	fmt.Println("Top down approach: ", dfs(n-1, sum))
}

func bottomUp(n, sum int) {
	var memo = make([][]bool, n)
	var v1, v2 bool
	for i := range memo {
		memo[i] = make([]bool, sum + 1)
	}

	for i := 0; i < n; i ++ {
		for j := 0; j < sum + 1; j++ {
			if j == 0 {
				memo[i][j] = true
				continue
			}

			// Handle negative index
			if i - 1 < 0 {
				v1 = false
			} else {
				v1 = memo[i-1][j]
			}

			// Handle negative index
			if j  - arr[i] < 0 || i - 1 < 0{
				v2 = false
			} else {
				v2 =  memo[i - 1][j  - arr[i]]
			}

			if j - arr[i] == 0 {
				v2 = true
			}

			if arr[i] <= j {
				memo[i][j] = v2 || v1
			} else {
				memo[i][j] = v1
			}
		}
	}
	fmt.Println("bottom up approach: ",memo[n-1][sum])
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