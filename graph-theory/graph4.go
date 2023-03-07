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
			if explore(iR, iC, grid, visited) == true {
				count ++
			}
		}
	} 
	fmt.Println("count: ", count)
}

func explore(iR, iC int, grid [][]string, visited map[string]bool) bool {
	if iC < 0 || iC >= len(grid[0]) {
		return false
	}

	if iR < 0 || iR >= len(grid) {
		return false
	}

	if grid[iR][iC] == "w" {
		return false
	}

	point := strconv.Itoa(iR) + "-" + strconv.Itoa(iC)

	if _, ok := visited[point]; ok {
		return false
	}

	visited[point] = true

	explore(iR - 1, iC, grid, visited)
	explore(iR + 1, iC, grid, visited)
	explore(iR, iC - 1, grid, visited)
	explore(iR, iC + 1, grid, visited)

	return true
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