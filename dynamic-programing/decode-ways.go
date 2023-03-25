package main

import "fmt"
import "strconv"

func main() {
	s := "121"
	way1(s)
}

func way1(s string) {
	var dfs func(string ) int
	dfs = func(str string) int {
		if str == "" {
			return 1
		}
		count := 0

		if len(str) >= 2 && isValid(str[0:2]) {
			count = count + dfs(str[2:])
		}
		if isValid(str[0:1]) {
			count = count + dfs(str[1:])
		}
		return count
	}
	fmt.Println(dfs(s))
}

func isValid(s string) bool {
	if s[0] == '0' {
		return false
	}
	val, err := strconv.Atoi(s)
    if err != nil {
        return false
    }

	if val >= 1 && val <= 26 {
		return true
	}
	return false
}