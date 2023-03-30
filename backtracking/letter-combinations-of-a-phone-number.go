package main 

import "fmt"

func main() {
	digits := "256"
	way1(digits)
}

func way1(digits string) {
	mapChar := make(map[byte]string)
	mapChar['2'] = "abc"
	mapChar['3'] = "def"
	mapChar['4'] = "ghi"
	mapChar['5'] = "jkl"
	mapChar['6'] = "mno"
	mapChar['7'] = "pqrs"
	mapChar['8'] = "tuv"
	mapChar['9'] = "wxyz"

	var dfs func(int)
	var res []string
	var part string = ""

	dfs = func(i int) {
		if i >= len(digits) {
			res = append(res, part)
			return
		}

		for _, c := range mapChar[digits[i]] {
			part = part + string(c)
			dfs(i+1)
			part = part[:len(part) - 1]
		}
	}
	dfs(0)
	fmt.Println(res)
}