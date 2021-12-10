package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	expected := 5934

	output := part1(input)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	expected := 26984457539

	output := part2(input)

	assert.Equal(t, expected, output)
}
