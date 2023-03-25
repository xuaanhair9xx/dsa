package main

import "fmt"

func main() {
	coins := []int{2}
	amount := 3
	way1(coins, amount)

}

func way1(coins []int, amount int) {
	var dfs func(int, int) 
	var min int = 1000

	dfs = func (total int, num int) {
		if total >= amount {
			if num < min && total == amount {
				min = num
			}
			return
		}

		for _, v := range coins {
			dfs(total + v, num + 1)
		}
	}
	dfs(0,0)
	fmt.Println(min)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}