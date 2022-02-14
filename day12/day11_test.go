package main

import (
	"testing"

	"github.com/pr00se/advent-of-code-2021/data"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `5483143223
	2745854711
	5264556173
	6141336146
	6357385478
	4167524645
	2176841721
	6882881134
	4846848554
	5283751526`

	expected := 1656

	grid, _ := data.ParseGrid(input)

	output := part1(grid)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := `5483143223
	2745854711
	5264556173
	6141336146
	6357385478
	4167524645
	2176841721
	6882881134
	4846848554
	5283751526`

	expected := 195

	grid, _ := data.ParseGrid(input)

	output := part2(grid)

	assert.Equal(t, expected, output)
}
