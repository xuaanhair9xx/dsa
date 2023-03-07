package main

import (
	"fmt"
	"strconv"
)
func isLandCount(grid [][]string) {
	visited := make(map[string]bool)
	count := 0
	for iR := 0; iR < len(grid); iR ++ {
		for iC := 0; iC < len(grid[0]); iC ++ {
			size := explore(iR, iC, grid, visited)
			if size != 0 {
				fmt.Println("size: ", size)
			}
		}
	} 
	fmt.Println("count: ", count)
}

func explore(iR, iC int, grid [][]string, visited map[string]bool) int {
	if iC < 0 || iC >= len(grid[0]) {
		return 0
	}

	if iR < 0 || iR >= len(grid) {
		return 0
	}

	if grid[iR][iC] == "w" {
		return 0
	}

	point := strconv.Itoa(iR) + "-" + strconv.Itoa(iC)

	if _, ok := visited[point]; ok {
		return 0
	}

	visited[point] = true
	size := 1

	size = size + explore(iR - 1, iC, grid, visited)
	size = size + explore(iR + 1, iC, grid, visited)
	size = size + explore(iR, iC - 1, grid, visited)
	size = size + explore(iR, iC + 1, grid, visited)

	return size
}

func main() {
	grid := [][]string {
		{"w","l","w","w","w"},
		{"w","l","w","w","w"},
		{"w","w","w","l","w"},
		{"w","w","l","l","w"},
		{"l","w","w","l","l"},
		{"l","l","w","w","w"},
	}
	isLandCount(grid);
}