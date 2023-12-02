package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("first/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Cubes in bag: 12 red cubes, 13 green cubes, and 14 blue cubes
	inputLines := strings.Split(string(content), "\n")
	sum := 0
	for _, v := range inputLines {
		parts := strings.Split(v, ":")

		gameNumber, err := strconv.Atoi(strings.Split(strings.TrimSpace(parts[0]), " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		games := strings.Split(parts[1], ";")

		possibleGame := true
		for _, game := range games {
			cubesOfColor := strings.Split(game, ",")
			for _, cube := range cubesOfColor {
				countAndColor := strings.Split(strings.TrimSpace(cube), " ")
				count, err := strconv.Atoi(countAndColor[0])
				if err != nil {
					log.Fatal(err)
				}

				color := countAndColor[1]
				if color == "red" && count > 12 {
					possibleGame = false
				}

				if color == "green" && count > 13 {
					possibleGame = false
				}

				if color == "blue" && count > 14 {
					possibleGame = false
				}
				fmt.Printf("%s %s\n", countAndColor[0], countAndColor[1])
			}
		}
		fmt.Printf("Games: %s %t\n", games, possibleGame)
		if possibleGame {
			sum += gameNumber
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
