package main

import (
	"os"
	"strings"
)

func main() {
	// Search for the * in the doc
	// find all numbers around the star
	// parse the numbers and multiply them, only if we found at least two numbers

	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(content), "\n")
	for rowIndex, row := range rows {
		for index, char := range row {
			if char != '*' {
				continue
			}

			topRow := min(rowIndex-1, 0)
			bottomRow := max(rowIndex+1, len(rows)-1)
			leftColumn := min(index-1, 0)
			rightColumn := max(index+1, len(row)-1)

		}
	}
}
