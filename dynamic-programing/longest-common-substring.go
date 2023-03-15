package main

import (
	"fmt"
)

func main() {
	str1 := "dwwGeeksmforGeeks"
	str2 := "eeeeeGeeksmforQuiz"
	brushForce(str1, str2)
	dynamicPrograme(str1, str2)
	dynamicPrograme2(str1, str2)
}

func brushForce(str1, str2 string) {
	max := 0
	for i := 0; i < len(str1); i++ {
		for j:= 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				flag := true
				k := i
				l := j
				count := 0
				for flag && l < len(str2) && k < len(str1) {
					if str1[k] == str2[l] {
						count++
						k++
						l++
					} else {
						flag = false
					}
				}
				if max < count {
					max = count
				}
			}
		}
	} 
}

func dynamicPrograme(str1, str2 string) {
	n := len(str1)

	var dp = make([][]int, n + 1)
	var max, x, y int
	max = 0
	for i := range dp {
		dp[i] = make([]int, len(str2) + 1)
	}

	for i := 0; i < n + 1; i ++ {
		for j := 0; j < len(str2) + 1; j++ {
			if j == 0 || i == 0 {
				continue
			}
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > max {
					max = dp[i][j]
					x = i - 1
					y = j - 1
				}
			}
		}
	} 
	for i := x + 1 - max ; i < x + 1 ; i ++ {
		fmt.Printf("%s", string(str1[i]))
	}
	fmt.Printf("\n max: %v pos-str1: %v pos-str2: %v", max, x, y)
}

func dynamicPrograme2(str1, str2 string) {
	n := len(str1)
	m := len(str2)
	var dp = make([][]int, n + 1)
	for i := range dp {
		dp[i] = make([]int, m + 1)
	}

	for i := 0; i < n + 1; i ++ {
		for j := 0; j < m + 1; j++ {
			if j == 0 || i == 0 {
				continue
			}
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	} 
	fmt.Println("Max: ", dp[n][m])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}