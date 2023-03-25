package main

import "fmt"

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	way1(s1,s2,s3)
}

func way1(s1, s2, s3 string) {
	var dfs func (int, int, string) bool

	dfs = func(i int, j int, str string) bool {

		if len(str) == len(s3) {
			if str == s3 {
				return true
			} else {
				return false
			}
		}

		check := false

		if len(s1[i:]) > 0 {
			check = check || dfs(i+1, j, str + s1[i:i+1])
		}
		if len(s2[j:]) > 0 {
			check = check || dfs(i, j + 1, str + s2[j:j+1])
		}
		return check
	}

	fmt.Println(dfs(0,0,""))
}