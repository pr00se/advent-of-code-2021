package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	var nums []int

	_, path, _, _ := runtime.Caller(0)
	path = filepath.Join(filepath.Dir(path), "input.txt")

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if len(s.Text()) > 0 {
			i, err := strconv.Atoi(s.Text())
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, i)
		}
	}

	fmt.Printf("Part 1 solution: %d\n", part1(nums))
	fmt.Printf("Part 2 solution: %d\n", part2(nums))

	fmt.Printf("Part 1 solution (generalized): %d\n", generalized(nums, 1))
	fmt.Printf("Part 2 solution (generalized): %d\n", generalized(nums, 3))
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
