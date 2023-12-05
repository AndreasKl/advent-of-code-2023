package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(content), "\n")
	seeds := parseSeeds(rows)
	maps := parseMappings(rows[2:])

	nearest := math.MaxInt
	for _, s := range seeds {
		nearest = min(maps.calculateLocation(s), nearest)
		fmt.Printf("\nLocation: %d\n\n", maps.calculateLocation(s))
	}
	fmt.Printf("Nearest: %d\n\n", nearest)

}

type overallMapping []mappings

func (m overallMapping) calculateLocation(toMap int) int {
	tmp := toMap
	for _, v := range m {
		tmp = v.mapWithMultiMappings(tmp)
	}
	return tmp
}

type mappings struct {
	name     string
	mappings []mapping
}

func (m mappings) mapWithMultiMappings(toMap int) int {
	tmp := toMap

	for _, v := range m.mappings {
		start := v.inboundStart
		end := v.inboundStart + v.inboundLength - 1
		if start <= tmp && end >= tmp {
			tmp = v.outboundStart + (-1 * (start - tmp))
			break
		}
	}
	return tmp
}

type mapping struct {
	inboundStart  int
	inboundLength int
	outboundStart int
}

func parseMappings(rows []string) overallMapping {
	var mm []mappings

	m := mappings{}
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		if unicode.IsLetter(rune(row[0])) {
			if m.name != "" {
				mm = append(mm, m)
			}
			m = mappings{}
			m.name = row
			continue
		}

		if unicode.IsDigit(rune(row[0])) {
			possibleNumbers := strings.Split(strings.TrimSpace(row), " ")
			ranges := stringSliceToIntSlice(possibleNumbers)
			m.mappings = append(m.mappings, mapping{
				outboundStart: ranges[0],
				inboundStart:  ranges[1],
				inboundLength: ranges[2],
			})
			continue
		}
	}

	mm = append(mm, m)

	return mm
}

func parseSeeds(rows []string) []int {
	seedsRaw :=
		strings.Split(
			strings.TrimSpace(
				strings.Split(rows[0], ":")[1],
			),
			" ",
		)

	return stringSliceToIntSlice(seedsRaw)
}

func stringSliceToIntSlice(stringSlice []string) []int {
	var intSlice []int
	for _, s := range stringSlice {
		seed, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, seed)
	}
	return intSlice
}
