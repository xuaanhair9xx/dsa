package main

import "fmt"
func main() {
	nums := []int {2, 3, -2, 4, -1}
	way1(nums)
}

func way1(nums []int) {
	var maxV int = 0
	var dfs func(int, int) int
	var n int = len(nums)

	dfs = func(i int, cur int) int {
		if i >= n {
			return 0
		}
		return max(cur * nums[i], dfs(i+1, cur * nums[i]))
	}

	for id, _ := range nums {
		maxV = max(maxV, dfs(id, 1))
	}

	fmt.Println(maxV)
}

func max(a, b int) int {
	if a > b {
		return a 
	} else {
		return b
	}
}