package data

import (
	"strconv"
	"strings"
)

// Grid represents a 2D plane of values
type Grid map[Point]int

// Point represents a point in 2D space
type Point struct {
	X, Y int
}

func (p Point) up() Point {
	return Point{X: p.X, Y: p.Y + 1}
}

func (p Point) down() Point {
	return Point{X: p.X, Y: p.Y - 1}
}

func (p Point) left() Point {
	return Point{X: p.X - 1, Y: p.Y}
}

func (p Point) right() Point {
	return Point{X: p.X + 1, Y: p.Y}
}

func (p Point) upRight() Point {
	return Point{X: p.X + 1, Y: p.Y + 1}
}

func (p Point) downRight() Point {
	return Point{X: p.X + 1, Y: p.Y - 1}
}

func (p Point) downLeft() Point {
	return Point{X: p.X - 1, Y: p.Y - 1}
}

func (p Point) upLeft() Point {
	return Point{X: p.X - 1, Y: p.Y + 1}
}

// Neighbors returns the 8 Points adjacent and diagonal to p
func (p Point) Neighbors() []Point {
	return append(p.Adjacent(), p.Diagonal()...)
}

// Adjacent returns the 4 Points adjacent to p
func (p Point) Adjacent() []Point {
	return []Point{p.up(), p.down(), p.left(), p.right()}
}

// Diagonal returns the 4 points diagonally adjacent to p
func (p Point) Diagonal() []Point {
	return []Point{p.upRight(), p.downRight(), p.downLeft(), p.upLeft()}
}

// ParseGrid parses the input string into a Grid
func ParseGrid(input string) (Grid, error) {
	var grid = Grid{}

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for y, line := range lines {
		for x, p := range strings.TrimSpace(line) {
			i, err := strconv.Atoi(string(p))
			if err != nil {
				return nil, err
			}

			grid[Point{X: x, Y: y}] = i
		}
	}

	return grid, nil
}

// ParsePoint parses a comma-separated coordinate pair into a Point
func ParsePoint(input string) (Point, error) {
	nums, err := ParseCommaSeparatedInts(input)
	if err != nil {
		return Point{}, err
	}

	return Point{X: nums[0], Y: nums[1]}, nil
}
