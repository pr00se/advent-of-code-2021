package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	expected := 198

	output := part1(input)

	if output != expected {
		t.Errorf("Expected: %d Got: %d", expected, output)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	expected := 230

	output := part2(input)

	if output != expected {
		t.Errorf("Expected: %d Got: %d", expected, output)
	}
}
