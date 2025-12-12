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

	beams := []Beam{{row: startRow, col: startCol}}
	splitCount := 0
	visited := make(map[Beam]bool)

	for len(beams) > 0 {
		currentBeam := beams[0]
		beams = beams[1:]

		// Move downward until a splitter or end of manifold
		for r := currentBeam.row + 1; r < len(manifold); r++ {
			if r < 0 || r >= len(manifold) || currentBeam.col < 0 || currentBeam.col >= len(manifold[r]) {
				// Beam exited the manifold
				break
			}

			if visited[Beam{row: r, col: currentBeam.col}] {
				// Already processed this path
				break
			}
			visited[Beam{row: r, col: currentBeam.col}] = true

			if manifold[r][currentBeam.col] == '^' {
				splitCount++
				// Splitter encountered, create two new beams
				// Left beam
				if currentBeam.col-1 >= 0 {
					beams = append(beams, Beam{row: r, col: currentBeam.col - 1})
				}
				// Right beam
				if currentBeam.col+1 < len(manifold[r]) {
					beams = append(beams, Beam{row: r, col: currentBeam.col + 1})
				}
				break // Current beam stops
			}
		}
	}

	fmt.Println(splitCount)
}
