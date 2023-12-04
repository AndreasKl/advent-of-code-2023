package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculatePointsFromPileOfCards(t *testing.T) {
	content, err := os.ReadFile("./../example.txt")
	require.NoError(t, err)

	scratchCardInputLine := strings.Split(string(content), "\n")
	pileOfCardsPoints := calculatePointsFromPileOfCards(scratchCardInputLine)

	assert.Equal(t, 13, pileOfCardsPoints)
}
