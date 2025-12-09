package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func isInvalidID(id int) bool {
	s := strconv.Itoa(id)

	if len(s) > 1 && s[0] == '0' {
		return false
	}

	n := len(s)
	if n%2 != 0 {
		return false 
	}

	halfLen := n / 2
	firstHalf := s[:halfLen]
	secondHalf := s[halfLen:]

	return firstHalf == secondHalf
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input.txt:", err)
		return
	}

	input := strings.TrimSpace(string(content))
	ranges := strings.Split(input, ",")

	var totalInvalidIDsSum int
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			fmt.Println("Invalid range format:", r)
			continue
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error parsing start ID:", parts[0], err)
			continue
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error parsing end ID:", parts[1], err)
			continue
		}

		for i := start; i <= end; i++ {
			if isInvalidID(i) {
				totalInvalidIDsSum += i
			}
		}
	}

	fmt.Println(totalInvalidIDsSum)
}
