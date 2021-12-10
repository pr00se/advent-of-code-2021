package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

type grid map[point]int

type point struct {
	x, y int
}

func (p point) up() point {
	return point{x: p.x, y: p.y + 1}
}

func (p point) down() point {
	return point{x: p.x, y: p.y - 1}
}

func (p point) left() point {
	return point{x: p.x - 1, y: p.y}
}

func (p point) right() point {
	return point{x: p.x + 1, y: p.y}
}

func (p point) neighbors() []point {
	return []point{p.up(), p.down(), p.left(), p.right()}
}

// parseInput parses the input string and returns the grid
func parseInput(input string) (grid, error) {
	var cave = grid{}

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for y, line := range lines {
		for x, p := range strings.TrimSpace(line) {
			i, err := strconv.Atoi(string(p))
			if err != nil {
				return nil, err
			}

			cave[point{x: x, y: y}] = i
		}
	}

	return cave, nil
}

// findLowPoints finds all the low points in the cave
func findLowPoints(cave grid) []point {
	var lowPoints []point

MainLoop:
	for p, d := range cave {
		// if height at this point is greater than at any of its neighbors,
		// this isn't the lowest point
		for _, n := range p.neighbors() {
			if x, ok := cave[n]; ok && d >= x {
				continue MainLoop
			}
		}

		lowPoints = append(lowPoints, p)
	}

	return lowPoints
}

// findBasin returns all the points in the same basin as p
func findBasin(cave grid, p point) []point {
	var (
		basin   []point
		toVisit []point
		visited = map[point]bool{}
	)

	visited[p] = true
	toVisit = p.neighbors()

	for len(toVisit) > 0 {
		visiting := toVisit
		toVisit = nil

		for _, v := range visiting {
			if d, ok := cave[v]; !ok || d == 9 {
				continue
			}

			visited[v] = true

			for _, n := range v.neighbors() {
				if !visited[n] {
					toVisit = append(toVisit, n)
				}
			}
		}
	}

	for p := range visited {
		basin = append(basin, p)
	}

	return basin
}

func part1(cave grid) int {
	var ratings int

	lowPoints := findLowPoints(cave)
	for _, p := range lowPoints {
		ratings += (1 + cave[p])
	}

	return ratings
}

func part2(cave grid) int {
	var sizes []int

	lowPoints := findLowPoints(cave)
	for _, p := range lowPoints {
		basin := findBasin(cave, p)
		sizes = append(sizes, len(basin))
	}

	sort.Ints(sizes)

	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	cave, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(cave))
	fmt.Printf("Part 2: %d\n", part2(cave))
}
