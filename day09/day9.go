package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/pr00se/advent-of-code-2021/data"
)

// findLowPoints finds all the low points in the cave
func findLowPoints(cave data.Grid) []data.Point {
	var lowPoints []data.Point

MainLoop:
	for p, h := range cave {
		// if height at this point is greater than at any of its adjacent
		// neighbors, this isn't the lowest point
		for _, n := range p.Adjacent() {
			if nh, ok := cave[n]; ok && h >= nh {
				continue MainLoop
			}
		}

		lowPoints = append(lowPoints, p)
	}

	return lowPoints
}

// findBasin returns all the points in the same basin as p
func findBasin(cave data.Grid, p data.Point) []data.Point {
	var (
		basin   []data.Point
		visited = map[data.Point]bool{}
	)

	visited[p] = true

	for visiting, toVisit := p.Adjacent(), []data.Point{}; len(visiting) > 0; visiting, toVisit = toVisit, nil {
		for _, v := range visiting {
			if h, ok := cave[v]; !ok || h == 9 {
				continue
			}

			visited[v] = true

			for _, n := range v.Adjacent() {
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

func part1(cave data.Grid) int {
	var ratings int

	lowPoints := findLowPoints(cave)
	for _, p := range lowPoints {
		ratings += (1 + cave[p])
	}

	return ratings
}

func part2(cave data.Grid) int {
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

	grid, err := data.ParseGrid(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(grid))
	fmt.Printf("Part 2: %d\n", part2(grid))
}
