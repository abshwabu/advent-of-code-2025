package main

import "fmt"

func countNeighbors(grid []string, row, col int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue // Skip the current cell
			}

			nRow, nCol := row+i, col+j
			if nRow >= 0 && nRow < rows && nCol >= 0 && nCol < cols {
				if grid[nRow][nCol] == '@' {
					count++
				}
			}
		}
	}
	return count
}

func ForkLift(lines []string) {
	accessible := 0
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			if lines[r][c] == '@' {
				if countNeighbors(lines, r, c) < 4 {
					accessible++
				}
			}
		}
	}
	fmt.Println(accessible)
}
