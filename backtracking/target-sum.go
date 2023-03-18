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
	visited = make([]bool, n)
	count = 0
	dfs(1, 0)
	
	fmt.Println(count)
}

func dfs(target int, pos int) {
	if pos >= n {
		if target == 0 {
			count++
		}
		return
	}

	for _, symbol := range "+-" {
		if !visited[pos] {
			visited[pos] = true
			if symbol == '+' {
				target = target + nums[pos]
			} else {
				target = target - nums[pos]
			}
			dfs(target, pos+1)
			visited[pos] = false
			if symbol == '+' {
				target = target - nums[pos]
			} else {
				target = target + nums[pos]
			}
		}
	}
}