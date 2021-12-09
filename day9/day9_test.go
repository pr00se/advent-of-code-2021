package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	input := `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`

	cave, err := parseInput(input)

	assert.Nil(err, "error should be nil")
	assert.Equal(50, len(cave), "should have parsed 50 points")
}

func TestPart1(t *testing.T) {
	input := `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`

	expected := 15

	cave, _ := parseInput(input)

	output := part1(cave)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`

	expected := 1134

	cave, _ := parseInput(input)

	output := part2(cave)

	assert.Equal(t, expected, output)
}
