package main

import (
	"fmt"
)
var arr [10]int
var n int = 3
var used map[int]bool

func main() {
	used := make(map[int]bool)
	for i := 0; i < n; i++ {
		used[i] = false
	} 
	permutation(0, used)
}

func show() {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	} 
	fmt.Println()
}

func permutation(pos int, used map[int]bool) {
	if (pos == n) {
		show()
	}
	for i := 0; i < n; i++ {		
		if used[i] == false {
			arr[pos] = i + 1
			used[i] = true
			permutation(pos + 1,used)
			used[i] = false
		}
	}
}