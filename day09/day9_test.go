package main

import (
	"testing"

	"github.com/pr00se/advent-of-code-2021/data"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`

	expected := 15

	grid, _ := data.ParseGrid(input)

	output := part1(grid)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`

	expected := 1134

	grid, _ := data.ParseGrid(input)

	output := part2(grid)

	assert.Equal(t, expected, output)
}
