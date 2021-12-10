package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

func parseInput(input string) ([]int, error) {
	var nums []int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, i)
	}

	return nums, nil
}

func countIncreases(nums []int, wsize int) int {
	var incs int

	// Samples that overlap between windows cancel each other out, so
	// we just need to compare the samples that are unique to each window
	for i := 0; i < len(nums)-wsize; i++ {
		if nums[i+wsize] > nums[i] {
			incs++
		}
	}

	return incs
}

func part1(nums []int) int {
	return countIncreases(nums, 1)
}

func part2(nums []int) int {
	return countIncreases(nums, 3)
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	nums, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(nums))
	fmt.Printf("Part 2: %d\n", part2(nums))
}
