package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scratchCardInputLines := strings.Split(string(content), "\n")

	var scratchCards scratchCards
	for _, inputLine := range scratchCardInputLines {
		scratchCards = append(scratchCards, parse(inputLine))
	}

	for index, c := range scratchCards {
		if c.points() == 0 {
			continue
		}

		start := keepInBounds(index+1, len(scratchCards))
		end := keepInBounds(index+1+c.points(), len(scratchCards))

		toIncrease := scratchCards[start:end]
		for i := 0; i < c.cardCount; i++ {
			for _, v := range toIncrease {
				v.increaseCardCount()
			}
		}
	}

	fmt.Printf("Overall card count: %d\n", scratchCards.overallCardCount())
}

func keepInBounds(index int, bound int) int {
	if index > bound {
		return bound
	}
	return index
}

type scratchCards []*scratchCard

func (s scratchCards) overallCardCount() int {
	sum := 0
	for _, c := range s {
		sum += c.cardCount
	}
	return sum
}

type scratchCard struct {
	game           string
	winningNumbers []int
	numbers        []int
	cardCount      int
}

func (s *scratchCard) points() int {
	count := 0
	for _, number := range s.numbers {
		for _, winningNumber := range s.winningNumbers {
			if number == winningNumber {
				count++
			}

		}
	}
	return count
}

func (s *scratchCard) increaseCardCount() {
	s.cardCount++
}

func parse(scratchCardInputLine string) *scratchCard {
	gameAndNumbers := strings.Split(scratchCardInputLine, ":")
	game := gameAndNumbers[0]
	numbersAndWinningNumbers := strings.Split(gameAndNumbers[1], "|")

	winningNumbers := strings.TrimSpace(numbersAndWinningNumbers[0])
	numbers := strings.TrimSpace(numbersAndWinningNumbers[1])

	return &scratchCard{
		game:           strings.TrimSpace(game),
		winningNumbers: parseNumbers(winningNumbers),
		numbers:        parseNumbers(numbers),
		cardCount:      1,
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
