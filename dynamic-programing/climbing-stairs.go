package main

import(
	"fmt"
)

func main() {
	n := 3
	var dfs func(int) int
	dfs = func(i int) int {
		if i >= n {
			if i == n {
				return 1
			} else {
				return 0
			}
		}
		return dfs(i+1) + dfs(i+2)
	}
	fmt.Println(dfs(0))
}