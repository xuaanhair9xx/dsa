package main

import "fmt"

func main() {
	nums := []int {1, 1, 2}
	way1(nums)
}

func way1(nums []int) {
	var dfs func(int)
	hashTable := make(map[int]int)
	n := len(nums)
	arr := make([]int, n)
	for _, v := range nums {
		if _, ok := hashTable[v]; ok {
			hashTable[v] = hashTable[v] + 1
		} else {
			hashTable[v] = 1
		}
	}
	dfs = func(pos int) {
		if pos == n {
			fmt.Println(arr)
			return
		}
		for value, count := range hashTable {
			if count > 0 {
				arr[pos] = value
				hashTable[value] = hashTable[value] - 1
				dfs(pos+1)
				hashTable[value] = hashTable[value] + 1
			}
		}
	}
	
	dfs(0)
}