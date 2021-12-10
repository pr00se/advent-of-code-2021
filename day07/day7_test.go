package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	expected := 37

	output := part1(input)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	expected := 168

	output := part2(input)

	assert.Equal(t, expected, output)
}
