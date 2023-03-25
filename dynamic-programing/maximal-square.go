package main

import "fmt"

func main() {
	matrix := [][]int {
		{1,1,1,1,1},
		{1,1,1,1,1},
		{1,1,1,1,1},
		{1,1,1,1,0},
	}
	way1(matrix)
}

func way1(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n)
	for i, _ := range matrix {
		dp[i] = make([]int, m)
	}

	k := n - 1
	l := m - 1

	for k >= 0 {
		for i := k; i >= 0; i -- {
			if i + 1 >= n {
				dp[i][l] = matrix[i][l]
				continue
			}
			if l + 1 >= m {
				dp[i][l] = matrix[i][l]
				continue
			}
			
			if matrix[i][l] == 0 || matrix[i+1][l] == 0 || matrix[i][l+1] == 0 || matrix[i+1][l+1] == 0 {
				dp[i][l] = matrix[i][l]
				continue
			}
			dp[i][l] = min(dp[i+1][l], min(dp[i][l+1], dp[i+1][l+1])) + 1
		} 
		for j := l - 1; j >= 0; j -- {
			if k + 1 >= n {
				dp[k][j] = matrix[k][j]
				continue
			}

			if matrix[k][j] == 0 || matrix[k+1][j] == 0 || matrix[k][j+1] == 0 || matrix[k+1][j+1] == 0 {
				dp[k][j] = matrix[k][j]
				continue
			}

			dp[k][j] = min(dp[k+1][j], min(dp[k][j+1], dp[k+1][j+1])) + 1
		} 
		k = k - 1
		l = l - 1
	}
	// [3][4] -> [2][4] -> [1][4] -> [0][4]
	// [3][3] -> [3][2] -> [3][1] -> [3][0]
	// [2][3] -> [1][3] -> [0][3]
	// [2][3] -> [2][1] -> [2][0]
	// [1][2] -> [0][2]
	// [1][1] -> [1][0]
	// [0][1]
	// [0][0]

	for _, f := range dp {
		fmt.Println(f)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}