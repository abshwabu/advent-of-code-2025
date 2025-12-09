package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isInvalidID checks if a number is an invalid ID.
// An ID is invalid if it's made only of some sequence of digits repeated twice.
// None of the numbers have leading zeroes; 0101 isn't an ID at all.
func isInvalidID(num int) bool {
	s := strconv.Itoa(num)

	// Check for leading zeros. A number like "0" is valid, but "01" is not.
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	// Iterate through possible substring lengths.
	// The repeating substring length 'subLen' must be a divisor of the total string length.
	// And 'subLen' must be less than the total string length.
	for subLen := 1; subLen < len(s); subLen++ {
		if len(s)%subLen == 0 { // Check if 'subLen' is a divisor
			sub := s[:subLen] // Get the potential repeating substring
			
			// Construct the string by repeating 'sub' (len(s) / subLen) times
			repeatedString := strings.Repeat(sub, len(s)/subLen)
			
			if s == repeatedString {
				return true // Found a repeating pattern
			}
		}
	}
	return false // No repeating pattern found
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input.txt: %v\n", err)
		return
	}

	line := strings.TrimSpace(string(content))
	rangesStr := strings.Split(line, ",")
	var totalSum int64 = 0

	for _, rStr := range rangesStr {
		parts := strings.Split(rStr, "-")
		if len(parts) != 2 {
			fmt.Printf("Invalid range format: %s\n", rStr)
			continue
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Invalid start number in range %s: %v\n", rStr, err)
			continue
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Invalid end number in range %s: %v\n", rStr, err)
			continue
		}

		for i := start; i <= end; i++ {
			if isInvalidID(i) {
				totalSum += int64(i)
			}
		}
	}

	fmt.Println(totalSum)
}