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
	profit := []int {2, 1, 4, 3}
	weight := []int {3, 3, 4, 2}
	wMax := 6
	recursive(profit, weight, wMax)
	topDown(profit, weight, wMax)
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
	memo := make(map[int]int)

	var dfs func(int, int) int
	dfs = func(i int, wMax int) int {
		if _, ok := memo[i]; ok {
			return memo[i]
		}
		if i < 0 {
			return 0
		}
		if weight[i] > wMax {
			memo[i] = dfs(i - 1, wMax)
		} else {
			memo[i] = max(profit[i] + dfs(i-1, wMax - weight[i]), dfs(i-1, wMax))
		}
		return memo[i]
	} 
	fmt.Println("Top down approach: ", dfs(n-1, wMax))
}