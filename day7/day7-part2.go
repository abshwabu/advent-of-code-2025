package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Beam struct {
	row int
	col int
}

// memoization table
var memo map[Beam]int

// countTimelines calculates the number of distinct timelines that can complete their journey
// starting from the given (row, col) position and moving downwards.
func countTimelines(r, c int, manifold []string) int {
	maxRows := len(manifold)
	maxCols := len(manifold[0]) // Assuming all rows have the same length

	// Base Cases:
	// If the current column is out of bounds, this path terminates without completing a journey.
	if c < 0 || c >= maxCols {
		return 0
	}
	// If the current row is at or beyond the manifold's bottom, this path has completed a journey.
	if r >= maxRows {
		return 1
	}

	// Memoization Check:
	currentBeam := Beam{row: r, col: c}
	if val, ok := memo[currentBeam]; ok {
		return val
	}

	totalPaths := 0
	
	// Simulate moving downwards from the current (r, c)
	// The particle moves straight down until it hits a splitter or the bottom.
	// We start checking for a splitter from the row *below* the current position.
	splitterRow := -1 // Stores the row where a splitter is found

	for i := r + 1; i < maxRows; i++ {
		// Check horizontal bounds at each step downwards
		if c < 0 || c >= maxCols { 
			break // Should not happen if `c` is valid at function entry
		}
		
		if manifold[i][c] == '^' {
			splitterRow = i
			break // Splitter found
		}
	}

	if splitterRow != -1 {
		// Splitter found at (splitterRow, c).
		// The current timeline splits into two.
		// These two new timelines effectively start their *downward journey* from
		// (splitterRow, c - 1) and (splitterRow, c + 1).
		totalPaths = countTimelines(splitterRow, c - 1, manifold) + countTimelines(splitterRow, c + 1, manifold)
	} else {
		// No splitter found below (r, c). The path continues straight down to the bottom.
		// The number of paths from this point is 1 (this single straight path).
		totalPaths = 1
	}

	memo[currentBeam] = totalPaths
	return totalPaths
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var manifold []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		manifold = append(manifold, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	startRow, startCol := -1, -1
	for r, line := range manifold {
		if idx := strings.Index(line, "S"); idx != -1 {
			startRow = r
			startCol = idx
			break
		}
	}

	if startRow == -1 {
		fmt.Println("Starting point 'S' not found in the manifold.")
		return
	}

	memo = make(map[Beam]int) // Initialize memoization table

	// Start counting timelines from the initial position of 'S'.
	totalTimelines := countTimelines(startRow, startCol, manifold)
	fmt.Println(totalTimelines)
}
