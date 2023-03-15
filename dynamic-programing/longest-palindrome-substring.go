package main

import (
	"fmt"
)

func main() {
	str := "forgeeksskeegfor"
	brushForce(str)
	dynamicPrograme(str)
}

func brushForce(str string) {
	max := 0
	start := 0

	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			flag := true
			for k := 0; k < (j - i) / 2; k ++ {
				if str[i+k] != str[j-k] {
					flag = false
				}
			}
			if flag && max < (j - i)/2 {
				max = (j - i)/2
				start = i
			}
		}
	}
	for i := start; i <= max + start; i++ {
		fmt.Printf("%s", string(str[i]))
	}
	fmt.Println()
}


func dynamicPrograme(str string) {
	n := len(str)
	var dp = make([][]int, n)
	var start, end, max int
	max = 1
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i ++ {
		dp[i][i] = 1
	}

	for len := 1; len < n; len = len + 2 {
		for i := 0; i < n - 1 - len; i ++ {
			j := i + len
			if str[i] != str[j] {
				continue
			}
			if str[i] == str[j] && (dp[i+1][j-1] == 1 || j - i == 1) {
				dp[i][j] = 1
				if j - i > max {
					max = j - i
					start = i
					end = j
				}
			}
		}
	} 
	


	for i := start; i <= end; i++ {
		fmt.Printf("%s", string(str[i]))
	}

}