package main

import "fmt"

func main() {
	s := "mississippi"
	p := "mis*is*p*."
	way1(s, p)
}

func way1(s, p string) {
	
	var dfs func(int, int) bool 

	dfs = func(i, j int) bool {
		if i >= len(s)-1 && j >= len(p) - 1 {
			return true
		}

		if j >= len(p) - 1 {
			return false
		}

		match := s[i] == p[j] || p[j] == '.'

		if len(p) - 1 > j + 1 && p[j+1] == '*' {
			if match {
				return dfs(i+1, j)
			} else {
				return dfs(i, j+2)
			}
		}

		if match {
			return dfs(i+1, j + 1)
		}

		return false
	}

	fmt.Println(dfs(0,0))
}