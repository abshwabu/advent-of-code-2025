package main

import "fmt"

// countNeighbors counts the number of '@' neighbors for a given cell.
// This function is the same as in day4-part1.go
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

// ForkLiftPart2 simulates the iterative removal of accessible paper rolls.
func ForkLiftPart2(lines []string) {
	totalRemoved := 0
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	for {
		removedThisIteration := 0
		toBeRemoved := make([][2]int, 0)

		// Find all accessible rolls in the current grid state
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[r]); c++ {
				if grid[r][c] == '@' {
					if countNeighbors(lines, r, c) < 4 {
						toBeRemoved = append(toBeRemoved, [2]int{r, c})
					}
				}
			}
		}

		// If no rolls were found to be removed, break the loop
		if len(toBeRemoved) == 0 {
			break
		}

		// Remove the accessible rolls and update the grid
		for _, pos := range toBeRemoved {
			r, c := pos[0], pos[1]
			if grid[r][c] == '@' { // Double check in case it was already removed by another fork lift in the same iteration
				grid[r][c] = '.'
				removedThisIteration++
			}
		}

		totalRemoved += removedThisIteration

		// Convert the updated rune grid back to string slice for countNeighbors
		// This is crucial because countNeighbors expects []string
		updatedLines := make([]string, len(grid))
		for i, row := range grid {
			updatedLines[i] = string(row)
		}
		lines = updatedLines // Update the 'lines' variable for the next iteration
	}

	fmt.Println(totalRemoved)
}
