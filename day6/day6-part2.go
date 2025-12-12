package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	if len(lines) < 2 {
		fmt.Println("Input file must contain at least two lines (numbers and operators).")
		return
	}

	// The last non-empty line contains the operators
	operatorLine := ""
	for i := len(lines) - 1; i >= 0; i-- {
		if trimmed := strings.TrimSpace(lines[i]); trimmed != "" {
			operatorLine = trimmed
			lines = lines[:i] // Remove operator line from number lines
			break
		}
	}

	if operatorLine == "" {
		fmt.Println("No operator line found.")
		return
	}

	// Reverse the order of number lines to process top-to-bottom
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	// Split operator line by space to find problem boundaries
	
	// Determine column widths and start/end for each problem
	// This is tricky because numbers can be multi-digit and there can be variable spacing.
	// The example suggests that problems are separated by *columns consisting only of spaces*.
	// Let's re-evaluate how to determine problem boundaries and extract numbers.

	// A better approach might be to transpose the grid, then split by empty lines to get problems.
	// Then, within each problem block, read numbers bottom-up.

	// Transpose the grid
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, maxLen)
		for j := 0; j < maxLen; j++ {
			if j < len(line) {
				grid[i][j] = rune(line[j])
			} else {
				grid[i][j] = ' '
			}
		}
	}

	transposed := make([][]rune, maxLen)
	for j := 0; j < maxLen; j++ {
		transposed[j] = make([]rune, len(lines))
		for i := 0; i < len(lines); i++ {
			transposed[j][i] = grid[i][j]
		}
	}

	// Join transposed columns back into strings, including the operator line
	transposedStrings := make([]string, maxLen)
	for j := 0; j < maxLen; j++ {
		var sb strings.Builder
		for i := 0; i < len(lines); i++ {
			sb.WriteRune(transposed[j][i])
		}
			var operatorChar rune = ' '
			if j < len(operatorLine) {
				operatorChar = rune(operatorLine[j])
			}
			transposedStrings[j] = sb.String() + string(operatorChar) // Append operator character
	}
	
	// Now, problems are separated by columns of only spaces.
	// Process right-to-left.
	grandTotal := int64(0)
	currentProblemNumbers := []int64{}
	currentOperator := ' '
	
	for j := maxLen - 1; j >= 0; j-- {
		col := transposedStrings[j]
		colNumbers := strings.TrimSpace(col[:len(col)-1]) // Numbers part of the column
		colOperator := rune(col[len(col)-1])              // Operator part of the column

		if colNumbers == "" && colOperator == ' ' { // This is a separator column
			if len(currentProblemNumbers) > 0 && currentOperator != ' ' {
				// Solve the accumulated problem
				result := solveProblem(currentProblemNumbers, currentOperator)
				grandTotal += result
				currentProblemNumbers = []int64{} // Reset for next problem
				currentOperator = ' '
			}
		} else {
			// Extract number from the column if it's not empty
			if colNumbers != "" {
				num, err := strconv.ParseInt(reverseString(colNumbers), 10, 64)
				if err != nil {
					fmt.Printf("Error parsing number %s: %v\n", reverseString(colNumbers), err)
					return
				}
				currentProblemNumbers = append(currentProblemNumbers, num)
			}

			// The operator for the problem is at the bottom of its column
			if colOperator != ' ' && strings.ContainsAny(string(colOperator), "+*") {
				currentOperator = colOperator
			}
		}
	}

	// Handle the last problem if any
	if len(currentProblemNumbers) > 0 && currentOperator != ' ' {
		result := solveProblem(currentProblemNumbers, currentOperator)
		grandTotal += result
	}

	fmt.Println("Grand Total:", grandTotal)
}

func solveProblem(numbers []int64, operator rune) int64 {
	if len(numbers) == 0 {
		return 0
	}

	// Numbers are already collected bottom-up (least significant digit at bottom)
	// Example: 4 + 431 + 623. Numbers would be [4, 431, 623] in our slice.
	// The problem statement says "most significant digit at the top and the least significant digit at the bottom".
	// When we constructed numbers from `reverseString(colNumbers)`, it correctly made '321' from '123' if read top-to-bottom.
	// But the example reads '431' from top '4', '3', '1'. So, `reverseString` was actually correct.

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		switch operator {
		case '+':
			result += numbers[i]
		case '*':
			result *= numbers[i]
		default:
			// Should not happen if operator parsing is correct
			fmt.Printf("Unknown operator: %c\n", operator)
			return 0
		}
	}
	return result
}

// reverseString reverses a string. Used to build numbers from columns.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
