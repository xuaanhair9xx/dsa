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

var profit []int
var weight []int

func main() {
	profit = []int {2, 1, 4, 3}
	weight = []int {3, 3, 4, 2}
	n := len(profit) - 1
	wMax := 6
	fmt.Println(dfs(n, wMax))
}

func dfs(i, wMax int) int {
	if i < 0 {
		return 0
	}
	if weight[i] > wMax {
		return dfs(i - 1, wMax)
	} else {
		return max(profit[i] + dfs(i-1, wMax - weight[i]), dfs(i-1, wMax))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}