package main

import (
	"fmt"
)
var arr []int
var n int = 3
var used map[int]bool

func main() {
	used = make(map[int]bool, n)
	arr = make([]int, n)
	permutation(0)
	fmt.Println(222)
	fmt.Println(arr)
}

func permutation(pos int) {
	if (pos == n) {
		fmt.Println(arr)
	}
	for i := 0; i < n; i++ {		
		if used[i] == false {
			arr[pos] = i + 1
			used[i] = true
			permutation(pos + 1)
			used[i] = false
			arr[pos] = 0
		}
	}
}