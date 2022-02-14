package main

import (
	"fmt"
	"log"

	"github.com/pr00se/advent-of-code-2021/data"
)

// runStep runs a single step of the simulation and returns the number of flashes
// that occur
func runStep(grid data.Grid) int {
	var (
		toFlash []data.Point
		flashed = map[data.Point]bool{}
	)

	// increment energy of all octopuses
	for p, e := range grid {
		grid[p] = e + 1
		// those with energy > 9 will flash
		if e >= 9 {
			toFlash = append(toFlash, p)
		}
	}

	// execute the flashes, and the flashes those flashes cause, and so on...
	for len(toFlash) > 0 {
		flashing := toFlash
		toFlash = nil

		for _, p := range flashing {
			// octopuses only flash once per turn
			if flashed[p] {
				continue
			}

			// increment energy of all neighbors
			for _, n := range p.Neighbors() {
				if e, ok := grid[n]; ok {
					grid[n] = e + 1
					// if this neighbor has energy > 9 and hasn't flashed yet, add it
					// to the list
					if e >= 9 && !flashed[n] {
						toFlash = append(toFlash, n)
					}
				}
			}

			flashed[p] = true
		}
	}

	// reset energy of octopuses that flashed
	for p := range flashed {
		grid[p] = 0
	}

	return len(flashed)
}

func part1(grid data.Grid) int {
	var flashes int

	for i := 0; i < 100; i++ {
		flashes += runStep(grid)
	}

	return flashes
}

func part2(grid data.Grid) int {
	for i := 1; ; i++ {
		if flashes := runStep(grid); flashes == len(grid) {
			return i
		}
	}
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

	// re-parse the grid since part1 modifies it
	grid, err = data.ParseGrid(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 2: %d\n", part2(grid))
}
