package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	expected := 150

	output := part1(input)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := []command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}
	expected := 900

	output := part2(input)

	assert.Equal(t, expected, output)
}
