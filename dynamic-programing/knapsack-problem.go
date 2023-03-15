package main

import (
	"fmt"
)

/**
We are given N items where each item has some weight and profit associated with it. 
We are also given a bag with capacity W, 
[i.e., the bag can hold at most W weight in it]. 
The target is to put the items into the bag such that the sum of 
profits associated with them is the maximum possible. 
**/

func main() {
	profit := []int {2, 1, 4, 3, 7, 2, 9, 4, 5, 8, 1, 40}
	weight := []int {3, 3, 4, 2, 4, 8, 2, 4, 7, 2, 1, 5}
	wMax := 10
	recursive(profit, weight, wMax)
	topDown(profit, weight, wMax)
	bottomUp(profit, weight, wMax)
}

// Recursive approach
func recursive(profit, weight []int, wMax int) {
	n := len(profit)
	var dfs func(int, int) int
	dfs = func(i int, wMax int) int {
		if i < 0 {
			return 0
		}
		if weight[i] > wMax {
			return dfs(i - 1, wMax)
		} else {
			return max(profit[i] + dfs(i-1, wMax - weight[i]), dfs(i-1, wMax))
		}
	} 
	fmt.Println("Recursive approach: ", dfs(n-1, wMax))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Top down approach
func topDown(profit, weight []int, wMax int) {
	n := len(profit)
	var memo = make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, wMax + 1)
	}

	var dfs func(int, int) int
	dfs = func(i int, wMax int) int {
		if i < 0 {
			return 0
		}
		if memo[i][wMax] != 0 {
			return memo[i][wMax]
		}
		if weight[i] > wMax {
			memo[i][wMax] = dfs(i - 1, wMax)
		} else {
			memo[i][wMax] = max(profit[i] + dfs(i-1, wMax - weight[i]), dfs(i-1, wMax))
		}
		return memo[i][wMax]
	} 
	fmt.Println("Top down approach: ", dfs(n-1, wMax))
}

func bottomUp(profit, weight []int, wMax int) {
	n := len(profit)

	var memo = make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, wMax + 1)
	}
	
	var v1, v2 int
	for i := 0; i < n; i++ {
		for j := 0; j < wMax+1; j++ {
			if i - 1 < 0 {
				v1 = 0
			} else {
				v1 = memo[i-1][j]
			}
			if j-weight[i] < 0 || i - 1  < 0 {
				v2 = 0
			} else {
				v2 = memo[i-1][j-weight[i]]
			}

			if (weight[i] > j) {
				memo[i][j] = v1
			} else {
				memo[i][j] = max(profit[i] + v2, v1)
			}
		} 
	}
	fmt.Println("bottom up approach: ",memo[n-1][wMax])
}