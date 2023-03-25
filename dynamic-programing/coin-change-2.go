package main

import (
	"fmt"
)

func main() {
	amount := 6
	coins := []int {1, 2, 5}
	way1(coins, amount)
	way2(coins, amount)

}

func way1( coins []int, amount int) {
	var dfs func(int, int) int
	dfs = func(total int, prev int) int {
		count := 0
		if total == amount {
			return 1
		}
		if total > amount {
			return 0
		}
		for i := prev; i < len(coins); i++ {
			count = count + dfs(total + coins[i], i)
		}
		return count
	}
	fmt.Println(dfs(0,0))
}

func way2(coins []int, amount int) {
	var dfs func(int, int) int
	dfs = func(total int, i int) int {
		if total == amount {
			return 1
		}
		if total > amount {
			return 0
		}
		if i == len(coins) {
			return 0
		}
		return dfs(total + coins[i], i) + dfs(total, i + 1)
	}
	fmt.Println(dfs(0,0))
}