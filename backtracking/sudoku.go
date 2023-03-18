package main

import (
	"fmt"
)

var board [9][9]int
var markRow [9][9]bool
var markCol [9][9]bool
var markMatric [3][3][9]bool

func main() {
	board = [9][9]int {
		{5,3,0,0,7,0,0,0,0},
		{6,0,0,1,9,5,0,0,0},
		{0,9,8,0,0,0,0,6,0},
		{8,0,0,0,6,0,0,0,3},
		{4,0,0,8,0,3,0,0,1},
		{7,0,0,0,2,0,0,0,6},
		{0,6,0,0,0,0,2,8,0},
		{0,0,0,4,1,9,0,0,5},
		{0,0,0,0,8,0,0,0,0},
	}

	
	for i := 0; i < len(board); i++  {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != 0 {
				num := board[i][j]
				markRow[i][num-1] = true
				markCol[j][num-1] = true
				markMatric[i/3][j/3][num-1] = true
			}
		}
	}
	solver(0,0)
	fmt.Println()
	printBoard()

}

func solver(i, j int) {
	if i < 9 && j < 9 {
		if board[i][j] == 0 {
			for z := 1; z <= 9; z++ {
				if !markRow[i][z-1] && !markCol[j][z-1] && !markMatric[i/3][j/3][z-1] {
					markRow[i][z-1] = true
					markCol[j][z-1] = true
					markMatric[i/3][j/3][z-1] = true
					board[i][j] = z
					solver(i, j+1)
					markRow[i][z-1] = false
					markCol[j][z-1] = false
					markMatric[i/3][j/3][z-1] = false
					board[i][j] = 0
				}
			}
		} else {
			solver(i, j+1)
		}
	} else if (i < 9 && j >= 9) {
		solver(i+1, 0)
	} else {
		printBoard()
	}
}

func printBoard() {
	for _, i := range board {
		fmt.Println(i)
	}
}