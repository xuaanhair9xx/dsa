package main

import(
	"fmt"
)

func main() {
	arr := []int {1, 3, 4, 0, 5, 0, 9, 2}
	fmt.Println(longestPathSubproblem(arr));
	fmt.Println(recursive(arr, len(arr) - 1))
}

func recursive(arr []int, n int) int {	
	if n == 0 {
		return 1
	}
    var count int = 1
	if arr[n] > arr[n-1] {
		count = 1 + recursive(arr, n-1)
	}
	return max(count, recursive(arr, n-1))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
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