package main

import(
	"fmt"
)

func main() {
	arr := []int {1, 3, 4, 0, 5, 0, 9, 2, 15}
	//fmt.Println(longestPathSubproblem(arr));
	recursive(arr, len(arr))
	topDown(arr, len(arr))
}

func recursive(arr []int, n int) {	
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 1
		}
		if arr[i] > arr[i-1] {
			return 1 + dfs(i-1)
		}
		return dfs(i-1)	
	}
	fmt.Println("Recursive approach: ", dfs(n-1))
}

func longestPathSubproblem(arr []int) int {
	n := len(arr)
	l := make([]int,n)
	for i := 0; i < n; i++ {
		l[i] = 1
	}
	
	max := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if (arr[i] > arr[j]) {
				l[i] = l[j] + 1
				if (l[i] > max) {
					max = l[i]
				}
			}
		}
	}

	return max
}