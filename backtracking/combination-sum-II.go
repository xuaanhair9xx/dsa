package main 

import "fmt"
import "sort"

func main() {
	candidates := []int{10, 1, 2,7, 6, 1,2,2, 5}
	target := 8
	way1(candidates, target)
}

func way1(candidates []int, target int) {
	var dfs func(int, int)  
	var part []int 
	n := len(candidates)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	
	dfs = func(i int, total int) {
		if total == target {
			fmt.Println(part)
			return 
		}
		if i >= n {
			return
		}
		prev := -1
		for j := i; j < n; j++ {
			if candidates[j] == prev {
				continue
			}
			part = append(part, candidates[j])
			dfs(j+1, total + candidates[j] )
			part = part[:len(part) - 1]
			prev = candidates[j]
		}
	}
	dfs(0,0)
}
