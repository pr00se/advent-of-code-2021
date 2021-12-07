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

func part1(nums []int) int {
	incs := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			incs++
		}
	}

	return incs
}

func part2(nums []int) int {
	incs := 0

	// The 2nd and 3rd samples in the previous window are the same as the
	// 1st and 2nd samples in the current window, so we just need to
	// compare the samples that are unique to each window
	for i := 0; i < len(nums)-3; i++ {
		if nums[i+3] > nums[i] {
			incs++
		}
	}

	return incs
}

func generalized(nums []int, wsize int) int {
	incs := 0

	for i := 0; i < len(nums)-wsize; i++ {
		if nums[i+wsize] > nums[i] {
			incs++
		}
	}

	return incs
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

	fmt.Printf("Part 1 (generalized): %d\n", generalized(nums, 1))
	fmt.Printf("Part 2 (generalized): %d\n", generalized(nums, 3))
}
