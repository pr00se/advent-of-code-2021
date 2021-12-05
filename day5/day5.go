package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type segment struct {
	start, end point
}

type column map[int]int

type ventMap map[int]column

// readInput reads the file at path and returns the line segments therein
func readInput(path string) ([]segment, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parseInput(string(content))
}

// parseInput parses the input string and returns the line segments therein
func parseInput(input string) ([]segment, error) {
	var (
		segments []segment
	)

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

	nums := strings.Split(strings.TrimSpace(input), ",")
	p.x, err = strconv.Atoi(nums[0])
	if err != nil {
		return p, err
	}
	p.y, err = strconv.Atoi(nums[1])
	if err != nil {
		return p, err
	}
	return p, nil
}

// plotSegment plots s on chart
func plotSegment(chart ventMap, s segment) {
	xLength := s.end.x - s.start.x
	yLength := s.end.y - s.start.y

	length := int(math.Abs(float64(xLength)))
	if length == 0 {
		length = int(math.Abs(float64(yLength)))
	}

	xStep := xLength / length
	yStep := yLength / length

	// segments include their endpoints, so we need to plot length+1 total points
	for i, x, y := 0, s.start.x, s.start.y; i <= length; i, x, y = i+1, x+xStep, y+yStep {
		if chart[x] == nil {
			chart[x] = make(column)
		}
		chart[x][y]++
	}
}

func part1(segments []segment) int {
	chart := make(ventMap)

	for _, s := range segments {
		if s.start.x == s.end.x || s.start.y == s.end.y {
			plotSegment(chart, s)
		}
	}

	dangerZones := 0
	for _, col := range chart {
		for _, rating := range col {
			if rating >= 2 {
				dangerZones++
			}
		}
	}

	return dangerZones
}

func part2(segments []segment) int {
	chart := make(ventMap)

	for _, s := range segments {
		plotSegment(chart, s)
	}

	dangerZones := 0
	for _, col := range chart {
		for _, rating := range col {
			if rating >= 2 {
				dangerZones++
			}
		}
	}

	return dangerZones
}

func main() {
	_, path, _, _ := runtime.Caller(0)
	path = filepath.Join(filepath.Dir(path), "input.txt")

	segments, err := readInput(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(segments))
	fmt.Printf("Part 2: %d\n", part2(segments))
}
