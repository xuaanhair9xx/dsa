package main

import (
	"fmt"
)

func main() {
	str := "forgeeksskeegfor"
	longestPalSubstr(str)
}

func longestPalSubstr(str string) {
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

	fmt.Printf("\n %v %v", max, start)
}