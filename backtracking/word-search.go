package main

import "fmt"
import "strconv"

func main() {
	board := [][]string {
		{"a", "b", "c", "e"},
		{"s", "f", "c", "s"},
		{"a", "d", "e", "e"},
	}
	word := "abccedfs"
	way1(board, word)
}

func way1(board [][]string, word string) {
	var dfs func(int, int, int) bool
	visted := make(map[string]bool)
	n := len(board)
	m := len(board[0])
	dfs = func(i, j, pos int) bool {
		if pos == len(word) {
			return true
		}
		if visted[toStr(i,j)] == true || i < 0 || i >= n || j < 0 || j >= m || board[i][j] != string(word[pos]){
			return false
		}
		visted[toStr(i,j)] = true
		result := dfs(i+1, j, pos + 1) || dfs(i-1, j, pos + 1) || dfs(i, j + 1, pos + 1) || dfs(i, j - 1, pos + 1)
		visted[toStr(i,j)] = false
		return result
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				fmt.Println("Existed")
			}
		}
	}
}

func toStr(i int, j int) string {
	return strconv.Itoa(i) + "-" + strconv.Itoa(j)
}