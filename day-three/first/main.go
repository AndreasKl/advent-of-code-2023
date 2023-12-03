package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	contentRows := strings.Split(string(content), "\n")
	for rowNumber, row := range contentRows {
		numberCollector := ""
		startOfNumber := 0
		for columnNumber, char := range row {
			if unicode.IsDigit(char) {
				if numberCollector == "" {
					startOfNumber = columnNumber
				}
				numberCollector += string(char)
				if columnNumber < len(row)-1 {
					continue
				}
			}

			if numberCollector != "" {
				endOfNumber := columnNumber

				if shouldBeCounted(contentRows, len(row), rowNumber, startOfNumber, endOfNumber) {
					number, err := strconv.Atoi(numberCollector)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("%d\n", number)
					sum += number
				}

				numberCollector = ""
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func shouldBeCounted(contentRows []string, rowLength int, rowNumber int, startOfNumber int, endOfNumber int) bool {
	beginOfCheckRangeY := max(rowNumber-1, 0)
	endOfCheckRangeY := min(rowNumber+1, len(contentRows)-1)

	beginOfCheckRangeX := max(startOfNumber-1, 0)
	endOfCheckRangeX := min(endOfNumber, rowLength-1)

	for checkY := beginOfCheckRangeY; checkY <= endOfCheckRangeY; checkY++ {
		for checkX := beginOfCheckRangeX; checkX <= endOfCheckRangeX; checkX++ {
			charToValidate := rune(contentRows[checkY][checkX])
			if !unicode.IsDigit(charToValidate) && !unicode.IsLetter(charToValidate) && charToValidate != '.' {
				return true
			}
		}
	}
	return false
}
