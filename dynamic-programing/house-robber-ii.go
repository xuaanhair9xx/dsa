package main

import "fmt"

func main() {
	nums := []int {2, 3, 3, 1, 9,1}
	way1(nums)
}

func way1(nums []int) {
	var dfs func(int, bool) int 
	n := len(nums)

	dfs = func(pos int, flag bool) int {
		var p1, p2 int

		if pos >= n {
			return 0
		} 

		if pos != n - 1|| !flag {
			if pos == 0 {
				flag = true
			}
			p1 = nums[pos] + dfs(pos + 2, flag)
		} 

		if pos + 1 != n - 1|| !flag {
			p2 = dfs(pos + 1, flag)
		}
		return min(p1, p2)
	}

	fmt.Println(dfs(0, false))
}

func min(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}