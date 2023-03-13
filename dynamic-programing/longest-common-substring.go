package main

import (
	"fmt"
)

func main() {
	str1 := "dwwGeeksmforGeeks"
	str2 := "eeeeeGeeksmQuiz"
	lcs(str1, str2)
}

func lcs(str1, str2 string) {
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

