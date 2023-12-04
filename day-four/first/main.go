package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scratchCardInputLine := strings.Split(string(content), "\n")
	pileOfCardsPoints := calculatePointsFromPileOfCards(scratchCardInputLine)

	fmt.Printf("Pile of cards points: %d\n", pileOfCardsPoints)
}

func calculatePointsFromPileOfCards(scratchCardInputLine []string) int {
	pileOfCardsPoints := 0
	for _, scratchCard := range scratchCardInputLine {
		sc := parse(scratchCard)
		pileOfCardsPoints += sc.points()
	}
	return pileOfCardsPoints
}

type scratchCard struct {
	game           string
	winningNumbers []int
	numbers        []int
}

func (s scratchCard) points() int {
	count := -1
	for _, number := range s.numbers {
		for _, winningNumber := range s.winningNumbers {
			if number == winningNumber {
				count++
			}

		}
	}
	return int(math.Pow(2.0, float64(count)))
}

func parse(scratchCardInputLine string) scratchCard {
	gameAndNumbers := strings.Split(scratchCardInputLine, ":")
	game := gameAndNumbers[0]
	numbersAndWinningNumbers := strings.Split(gameAndNumbers[1], "|")

	winningNumbers := strings.TrimSpace(numbersAndWinningNumbers[0])
	numbers := strings.TrimSpace(numbersAndWinningNumbers[1])

	return scratchCard{
		game:           strings.TrimSpace(game),
		winningNumbers: parseNumbers(winningNumbers),
		numbers:        parseNumbers(numbers),
	}

}

func parseNumbers(numbersText string) []int {
	numbers := strings.Split(numbersText, " ")
	var parsedNumbers []int
	for _, number := range numbers {
		trimmedNumber := strings.TrimSpace(number)
		if trimmedNumber == "" {
			continue
		}

		parsedNumber, err := strconv.Atoi(trimmedNumber)
		if err != nil {
			log.Fatal(err)
		}
		parsedNumbers = append(parsedNumbers, parsedNumber)
	}
	return parsedNumbers
}
