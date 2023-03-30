package main

import "fmt"
import "strconv"

func main() {
	s := "90890"
	way1(s)
}

func way1(s string) {
	var dfs func(int) 
	var isDescending func(int, int) bool
	var part []string 

	dfs = func(i int) {
		if i >= len(s) && len(part) > 1 {
			fmt.Println(part)
			fmt.Println("existed.")
			return
		}

		for j := i; j < len(s); j ++ {
			if isDescending(i, j + 1) {
				part = append(part, s[i:j+1])
				dfs(j+1)
				part = part[:len(part) - 1]
			} 
		}
	}

	isDescending = func(i, j int) bool {
		if i != j {
			if len(part) == 0 || cv(part[len(part)-1]) - cv(s[i:j]) == 1 {
				return true
			}
		} 
		return false
	}
	dfs(0)
}

func cv(str string) int {
	i, err := strconv.Atoi(str)
    if err != nil {
        panic(err)
    }
	return i
}

func way2(s string) {
	var dfs func(int, int)

	dfs = func(index, prev) {
		if index == len(s) {
			return true
		}

		for j := index: j < len(s); j++ {
			val := cv(s[index:j+1])
			if val + 1 == prev && dfs(j+1, val) {
				return true
			}
		}
		return false
	}

	for i := 0; i < len(s) - 1; i++ {
		val := cv(s[:i+1])
		if dfs(i+1, val) {
			fmt.Println("existed")
		}
	}
}