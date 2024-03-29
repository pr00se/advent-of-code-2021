package main

import (
	"fmt"
	"log"

	"github.com/pr00se/advent-of-code-2021/data"
)

// simulateFish simulates the growth of our school and returns the total number
// of fish alive after the given number of days
func simulateFish(fish []int, days int) int {
	newFish := make([]int, days)

	// project new fish generated by fish currently alive
	for _, f := range fish {
		newFish[f]++
		// more offspring every 7 days, forever
		n := f + 7
		for n < len(newFish) {
			newFish[n]++
			n += 7
		}
	}

	// step through each day, and project new fish generated by fish born on that day
	for i := 0; i < len(newFish); i++ {
		// first offspring after 9 days
		n := i + 9
		for n < len(newFish) {
			newFish[n] += newFish[i]
			// more offspring every 7 days, forever
			n += 7
		}
	}

	// return sum of all fish born on each day + original fish
	total := 0
	for _, count := range newFish {
		total += count
	}
	return total + len(fish)
}

func part1(fish []int) int {
	return simulateFish(fish, 80)
}

func part2(fish []int) int {
	return simulateFish(fish, 256)
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	fish, err := data.ParseCommaSeparatedInts(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(fish))
	fmt.Printf("Part 2: %d\n", part2(fish))
}
