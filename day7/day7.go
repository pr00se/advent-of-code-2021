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

// readInput reads the file at path and returns the crab positions therein
func readInput(path string) ([]int, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parseInput(string(content))
}

// parseInput parses the input string and returns the crab positions therein
func parseInput(input string) ([]int, error) {
	var (
		crabs []int
	)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		nums := strings.Split(strings.TrimSpace(line), ",")

		for _, n := range nums {
			i, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			crabs = append(crabs, i)
		}
	}

	return crabs, nil
}

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
	_, path, _, _ := runtime.Caller(0)
	path = filepath.Join(filepath.Dir(path), "input.txt")

	crabs, err := readInput(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(crabs))
	fmt.Printf("Part 2: %d\n", part2(crabs))
}
