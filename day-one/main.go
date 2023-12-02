package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	content, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	sum := 0
	currentLine := 1
	for _, line := range lines {
		if line == "" {
			continue
		}

		firstNumber, err := firstNumericCharFromSlice(line)
		if err != nil {
			fmt.Printf("Error on line %d: %s\n", currentLine, line)
			log.Fatal(err)
		}
		lastNumber, err := lastNumericCharFromSlice(line)
		if err != nil {
			fmt.Printf("Error on line %d: %s\n", currentLine, line)
			log.Fatal(err)
		}
		number, err := strconv.Atoi(fmt.Sprintf("%s%s", firstNumber, lastNumber))
		if err != nil {
			fmt.Printf("Error on line %d: %s\n", currentLine, line)
			log.Fatal(err)
		}
		sum += number
		currentLine++
	}
	fmt.Printf("Coordinates are %d!\n", sum)
}

func lastNumericCharFromSlice(line string) (string, error) {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			return string(line[i]), nil
		}
		soFar := line[i:]
		for k, v := range numbers {
			if strings.Contains(soFar, k) {
				return v, nil
			}
		}
	}
	return "", fmt.Errorf("no numeric char found")
}

func firstNumericCharFromSlice(line string) (string, error) {
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			return string(line[i]), nil
		}
		soFar := line[:i+1]
		for k, v := range numbers {
			if strings.Contains(soFar, k) {
				return v, nil
			}
		}
	}
	return "", fmt.Errorf("no numeric char found")
}
