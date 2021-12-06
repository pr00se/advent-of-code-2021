package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	input := `3,4,3,1,2`

	expectedFish := []int{3, 4, 3, 1, 2}

	fish, err := parseInput(input)

	assert.Nil(err, "error should be nil")
	assert.Equalf(expectedFish, fish, "fish should be parsed correctly")
}

func TestPart1(t *testing.T) {
	input := `3,4,3,1,2`

	expected := 5934

	fish, _ := parseInput(input)

	output := part1(fish)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := `3,4,3,1,2`

	expected := 26984457539

	fish, _ := parseInput(input)

	output := part2(fish)

	assert.Equal(t, expected, output)
}
