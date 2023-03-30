package main 

import "fmt"

func main() {
	s := "aab"
	way1(s)
}

func way1(s string) {
	var dfs func(int)
	var isPali func(int, int) bool

	var res [][]string
	var part []string

	dfs = func(i int) {
		if i >= len(s) {
			temp := make([]string, len(part))
			copy(temp, part)
			res = append(res, temp)
			return
		}

		for j := i; j < len(s); j++ {
			if isPali(i, j) {
				part = append(part, s[i:j+1])
				dfs(j+1)
				part = part[:len(part) - 1]
			}
		}
	}

	isPali = func(l, r int) bool {
		for l <= r {
			if s[l] != s[r] {
				return false
			}
			l = l + 1
			r = r - 1
		}
		return true
	}

	dfs(0)
	fmt.Println(res)
}