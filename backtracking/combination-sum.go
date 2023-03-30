package main

import "fmt"

func main() {
	candidates := []int {2, 3, 6, 7, 8}
	target := 8
	way1(candidates, target)
	way2(candidates, target)
}

func way1(candidates []int, target int) {
	var dfs func(int, int)
	var res [][]int
	var part []int

	dfs = func(i int, total int) {
		if total >= target {
			if total == target {
				temp := make([]int, len(part))
				copy(temp, part)
				res = append(res, temp)
			}
			return
		}

		for j := i; j < len(candidates); j++ {
			total = total + candidates[j]
			part = append(part, candidates[j])
			dfs(j, total)
			total = total - candidates[j]
			part = part[:len(part) - 1]
		}
	}
	dfs(0,0)
	fmt.Println(res)
}

func way2(candidates []int, target int) {
	var res [][]int
	var part []int
	var dfs func(int, int)

	dfs = func(i, total int) {
		if total == target {
			temp := make([]int, len(part))
			copy(temp, part)
			res = append(res, temp)
			return
		}

		if i >= len(candidates) || total > target {
			return
		}

		part = append(part, candidates[i])
		dfs(i, total + candidates[i])
		part = part[:len(part) - 1]
		dfs(i+1, total)
	}
	dfs(0,0)
	fmt.Println(res)
}