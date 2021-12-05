package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	input := `0,9 -> 5,9
	8,0 -> 0,8
	9,4 -> 3,4
	2,2 -> 2,1
	7,0 -> 7,4
	6,4 -> 2,0
	0,9 -> 2,9
	3,4 -> 1,4
	0,0 -> 8,8
	5,5 -> 8,2`

	segments, err := parseInput(input)

	assert.Nil(err, "error should be nil")
	assert.Equalf(10, len(segments), "should have parsed 10 segments")
}

func TestPart1(t *testing.T) {
	input := `0,9 -> 5,9
	8,0 -> 0,8
	9,4 -> 3,4
	2,2 -> 2,1
	7,0 -> 7,4
	6,4 -> 2,0
	0,9 -> 2,9
	3,4 -> 1,4
	0,0 -> 8,8
	5,5 -> 8,2`

	expected := 5

	segments, _ := parseInput(input)

	output := part1(segments)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := `0,9 -> 5,9
	8,0 -> 0,8
	9,4 -> 3,4
	2,2 -> 2,1
	7,0 -> 7,4
	6,4 -> 2,0
	0,9 -> 2,9
	3,4 -> 1,4
	0,0 -> 8,8
	5,5 -> 8,2`

	expected := 12

	segments, _ := parseInput(input)

	output := part2(segments)

	assert.Equal(t, expected, output)
}
