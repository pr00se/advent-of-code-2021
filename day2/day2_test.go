package main

import "testing"

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

	if output != expected {
		t.Errorf("Expected: %d Got: %d", expected, output)
	}
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

	if output != expected {
		t.Errorf("Expected: %d Got: %d", expected, output)
	}
}
