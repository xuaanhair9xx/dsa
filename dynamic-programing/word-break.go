package main

import(
	"fmt"
)

func main() {
	s := "leetleetcodeleet"
	wordDict := []string{"leet", "code"}
	way1(s, wordDict)
}

func way1(s string, wordDict []string) {

	var dfs func(int, int) bool 

	dfs = func(curPos, iWord int) bool {
		if curPos >= len(s) {
			return true
		}

		if iWord >= len(wordDict) {
			return false
		}

		check := true
		for i := 0; i < len(wordDict[iWord]); i++ {
			if wordDict[iWord][i] != s[curPos+i] {
				check = false
				break
			}
		}

		if check {
			curPos = curPos + len(wordDict[iWord])
			return dfs(curPos, 0)
		} else {
			return dfs(curPos, iWord + 1)
		}
	}

	fmt.Println(dfs(0,0))
}