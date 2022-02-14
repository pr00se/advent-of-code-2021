package main

import (
	"fmt"
	"log"
	"math"

	"github.com/pr00se/advent-of-code-2021/data"
)

type fuelFunc func(d int) int

// optimizeCrabFuel determines the minimum amount of fuel necessary to align
// crab subs, given their starting positions and a fuel consumption function
func optimizeCrabFuel(crabs []int, f fuelFunc) int {
	counts := map[int]int{}

	// determine starting position; assume all positions are positive
	startPos := -1
	for _, p := range crabs {
		counts[p]++
		if startPos == -1 || p < startPos {
			startPos = p
		}
	}

	fuel := -1
	// loop calculating fuel consumption at positions startPos..n, stopping once consumption
	// starts to increase
	for pos, nextFuel := startPos, 0; ; pos, fuel, nextFuel = pos+1, nextFuel, 0 {
		for p, c := range counts {
			dist := int(math.Abs(float64(p - pos)))
			nextFuel += f(dist) * c
		}

		if fuel != -1 && nextFuel > fuel {
			break
		}
	}

	return fuel
}

func part1(crabs []int) int {
	return optimizeCrabFuel(crabs, func(d int) int { return d })
}

func part2(crabs []int) int {
	// definitely didn't google a formula I'd forgotten since high school...
	return optimizeCrabFuel(crabs, func(d int) int { return ((d * d) + d) / 2 })
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	crabs, err := data.ParseCommaSeparatedInts(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(crabs))
	fmt.Printf("Part 2: %d\n", part2(crabs))
}
