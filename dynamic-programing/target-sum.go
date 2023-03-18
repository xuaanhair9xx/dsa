package main

import (
	"fmt"
)

/**
https://www.youtube.com/watch?v=g0npyaQtAQM
I want to build an expression out of nums by adding one of the symbols '+' and '-'
before each integer in nums and then concennate all the integers.
For example: if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and
concatenate them to build the expression "+2-1"
Return the number of different expressions that you can build, which evaluates to target.
Example 1: nums = [1,1,1,1,1], target = 3
Output: 5
*/

var nums []int
var count int
var n int
var visited []bool

func main() {
	nums = []int{1,1,1,1,1}
	n = len(nums)
	var dfs func(int, int) int
	dfs = func(i int, target int) int {
		if i >= n {
			if target == 0 {
				return 1
			} else {
				return 0
			}
		}
		return dfs(i + 1, target - nums[i]) + dfs(i+1, target + nums[i])
	}
	fmt.Println(dfs(0, 3))
}