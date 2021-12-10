package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

type point struct {
	x, y int
}

type segment struct {
	start, end point
}

type column map[int]int

type ventMap map[int]column

// parseInput parses the input string and returns the line segments therein
func parseInput(input string) ([]segment, error) {
	var segments []segment

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		pairs := strings.Fields(line)

		if len(pairs) != 3 {
			return nil, fmt.Errorf("invalid input: %s", line)
		}

		start, err := parsePoint(pairs[0])
		if err != nil {
			return nil, err
		}
		end, err := parsePoint(pairs[2])
		if err != nil {
			return nil, err
		}

		s := segment{
			start: start,
			end:   end,
		}
		segments = append(segments, s)
	}

	return segments, nil
}

// parsePoint parses a comma-separated coordinate pair into a point
func parsePoint(input string) (point, error) {
	var (
		p   point
		err error
	)

	nums, err := data.ParseCommaSeparatedInts(input)
	if err != nil {
		return p, err
	}
	p.x = nums[0]
	p.y = nums[1]

	return p, nil
}

// plotSegment plots s on chart and return the number of new danger zones created
func plotSegment(chart ventMap, s segment) int {
	xLength := s.end.x - s.start.x
	yLength := s.end.y - s.start.y

	length := int(math.Abs(float64(xLength)))
	if length == 0 {
		length = int(math.Abs(float64(yLength)))
	}

	xStep := xLength / length
	yStep := yLength / length

	dangerZones := 0

	// segments include their endpoints, so we need to plot length+1 total points
	for i, x, y := 0, s.start.x, s.start.y; i <= length; i, x, y = i+1, x+xStep, y+yStep {
		if chart[x] == nil {
			chart[x] = make(column)
		}
		chart[x][y]++

		if chart[x][y] == 2 {
			dangerZones++
		}
	}

	return dangerZones
}

func part1(segments []segment) int {
	chart := ventMap{}

	dangerZones := 0
	for _, s := range segments {
		if s.start.x == s.end.x || s.start.y == s.end.y {
			dangerZones += plotSegment(chart, s)
		}
	}

	return dangerZones
}

func part2(segments []segment) int {
	chart := ventMap{}

	dangerZones := 0
	for _, s := range segments {
		dangerZones += plotSegment(chart, s)
	}

	return dangerZones
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	segments, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(segments))
	fmt.Printf("Part 2: %d\n", part2(segments))
}
