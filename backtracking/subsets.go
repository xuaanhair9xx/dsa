package main

import "fmt"

func main() {
	nums := []int {1, 2, 3}
	way1(nums)
}

func way1(nums []int) {
	var dfs func(int, int)
	n := len(nums)
	dfs = func(pos int, bitMask int) {
		if pos == n {
			printResult(bitMask, nums)
			return
		}
		next := pos + 1
		dfs(next, bitMask)
		dfs(next, bitMask | 1 << next)
	}
	dfs(0, 1)
}

func printResult(bitMask int, nums []int) {
	var arr []int
	for i := 1; i <= len(nums); i ++ {
		if (1 << i) & bitMask != 0 {
			arr = append(arr, nums[i-1])
		}
	}
	fmt.Println(arr)
}