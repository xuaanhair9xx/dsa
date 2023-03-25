package main 

import "fmt"

func main() {
	costs := [][]int {{17,2,17},{16,16,5},{14,3,19}}
	way1(costs)
}

func way1(costs [][]int) {
	var dfs func(int, int) int
	n := len(costs)
	listColor := []int {0, 1, 2} 
	dfs = func(house int, color int) int {
		if house >= n {
			return 0
		}
		var otherColor []int
		for _, v := range listColor {
			if v != color {
				otherColor = append(otherColor, v)
			}
		}
		return min(costs[house][color] + dfs(house + 1, otherColor[0]), costs[house][color] + dfs(house + 1, otherColor[1]))
	}
	fmt.Println(min(dfs(0,0), min(dfs(0,1), dfs(0,2))))
}

func min(a, b int) int {
	if a > b {
		return b 
	} else {
		return a
	}
} 