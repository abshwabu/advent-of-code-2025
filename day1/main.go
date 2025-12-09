package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var start = 50
var zeroCounter = 0

func moves(move string) {
	n, _ := strconv.Atoi(move[1:])
	dir := 1
	if move[0] == 'L' || move[0] == 'l' {
		dir = -1
	}

	// -------- count zero hits during movement --------
	var first int
	if dir == 1 { // R
		first = (100 - start) % 100
	} else { // L
		first = start % 100
	}

	// first == 0 means starting on zero, but start doesn't count
	if first == 0 {
		first = 100
	}

	if n >= first {
		zeroCounter++                  // first
		zeroCounter += (n - first) / 100 // additional
	}

	// -------- move dial --------
	start = (start + dir*n%100 + 100) % 100
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	for scanner.Scan() {
		moves(scanner.Text())
	}

	fmt.Println(zeroCounter)
}

