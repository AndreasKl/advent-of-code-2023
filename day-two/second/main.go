package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("second/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputLines := strings.Split(string(content), "\n")
	overall := 0
	for _, v := range inputLines {
		parts := strings.Split(v, ":")

		games := strings.Split(parts[1], ";")
		minimumValidCubes := map[string]int{}
		for _, game := range games {
			cubesOfColor := strings.Split(game, ",")
			for _, cube := range cubesOfColor {
				countAndColor := strings.Split(strings.TrimSpace(cube), " ")
				count, err := strconv.Atoi(countAndColor[0])
				if err != nil {
					log.Fatal(err)
				}

				color := countAndColor[1]
				if _, ok := minimumValidCubes[color]; !ok {
					minimumValidCubes[color] = count
					continue
				}

				if count > minimumValidCubes[color] {
					minimumValidCubes[color] = count
				}
			}
		}

		sum := 0
		for _, count := range minimumValidCubes {
			if sum == 0 {
				sum = count
				continue
			}
			sum *= count
		}

		overall += sum
	}
	fmt.Printf("overall: %d\n", overall)
}
