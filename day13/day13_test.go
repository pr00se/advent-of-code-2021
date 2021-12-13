package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	input := `6,10
	0,14
	9,10
	0,3
	10,4
	4,11
	6,0
	6,12
	4,1
	0,13
	10,12
	3,4
	3,0
	8,4
	1,10
	2,14
	8,10
	9,0

	fold along y=7
	fold along x=5`

	grid, instructions, err := parseInput(input)

	assert.Nil(err, "error should be nil")
	assert.Equal(18, len(grid), "should have parsed 18 points")
	assert.Equal(2, len(instructions), "should have parsed 2 instructions")
}

func TestPart1(t *testing.T) {
	input := `6,10
	0,14
	9,10
	0,3
	10,4
	4,11
	6,0
	6,12
	4,1
	0,13
	10,12
	3,4
	3,0
	8,4
	1,10
	2,14
	8,10
	9,0

	fold along y=7
	fold along x=5`

	expected := 17

	grid, folds, _ := parseInput(input)

	output := part1(grid, folds)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := `6,10
	0,14
	9,10
	0,3
	10,4
	4,11
	6,0
	6,12
	4,1
	0,13
	10,12
	3,4
	3,0
	8,4
	1,10
	2,14
	8,10
	9,0

	fold along y=7
	fold along x=5`

	expected := "00000\n0   0\n0   0\n0   0\n00000\n"

	grid, folds, _ := parseInput(input)

	output := part2(grid, folds)

	assert.Equal(t, expected, output)
}
