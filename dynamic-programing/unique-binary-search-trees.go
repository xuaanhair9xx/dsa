package main

import "fmt"

func main() {
	n := 3
	way1(n)
}

func way1(n int) {
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 || i == 1 {
			return 1
		}
		total := 0
		for num := 0; num < i; num ++ {
			total = total + dfs(num) * dfs(i - 1 - num) 
		}
		return total
	}
	fmt.Println(dfs(n))
}