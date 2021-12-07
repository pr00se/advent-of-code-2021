package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7

	output := part1(input)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 5

	output := part2(input)

	assert.Equal(t, expected, output)
}

func TestGeneralized(t *testing.T) {
	tests := []struct {
		nums     []int
		wsize    int
		expected int
	}{
		{
			nums:     []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			wsize:    1,
			expected: 7,
		},
		{
			nums:     []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			wsize:    3,
			expected: 5,
		},
	}

	for _, tc := range tests {
		output := generalized(tc.nums, tc.wsize)
		assert.Equal(t, tc.expected, output)
	}
}
