package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	input := `16,1,2,0,4,2,7,1,2,14`

	expectedPositions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	fish, err := parseInput(input)

	assert.Nil(err, "error should be nil")
	assert.Equalf(expectedPositions, fish, "crab positions should be parsed correctly")
}

func TestPart1(t *testing.T) {
	input := `16,1,2,0,4,2,7,1,2,14`

	expected := 37

	crabs, _ := parseInput(input)

	output := part1(crabs)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := `16,1,2,0,4,2,7,1,2,14`

	expected := 168

	crabs, _ := parseInput(input)

	output := part2(crabs)

	assert.Equal(t, expected, output)
}
