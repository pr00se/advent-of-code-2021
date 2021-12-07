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
	counts := make(map[int]int)

	// determine fuel necessary at position 0
	fuel := 0
	for _, p := range crabs {
		counts[p]++
		fuel += f(p)
	}

	// loop calculating fuel consumption at positions 1..n, stopping once consumption
	// starts to increase again
	i, nextFuel := 1, 0
	for {
		for p, c := range counts {
			dist := int(math.Abs(float64(p - i)))
			nextFuel += f(dist) * c
		}

		if nextFuel > fuel {
			break
		}

		i++
		fuel = nextFuel
		nextFuel = 0
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
