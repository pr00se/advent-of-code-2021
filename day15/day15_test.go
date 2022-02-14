package main

import (
	"testing"

	"github.com/pr00se/advent-of-code-2021/data"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `1163751742
	1381373672
	2136511328
	3694931569
	7463417111
	1319128137
	1359912421
	3125421639
	1293138521
	2311944581`

	expected := 40

	grid, _ := data.ParseGrid(input)

	output := part1(grid)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := `1163751742
	1381373672
	2136511328
	3694931569
	7463417111
	1319128137
	1359912421
	3125421639
	1293138521
	2311944581`

	expected := 315

	grid, _ := data.ParseGrid(input)

	output := part2(grid)

	assert.Equal(t, expected, output)
}
