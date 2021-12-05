package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// readInput reads the file at path and returns the bingo calls and boards contained therein
func readInput(path string) ([]int, []*board, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	return parseInput(string(content))
}

// parseInput parses the input string and returns the bingo calls and boards contained therein
func parseInput(input string) ([]int, []*board, error) {
	var (
		calls  []int
		boards []*board
	)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Parse list of bingo calls
	nums := strings.Split(strings.TrimSpace(lines[0]), ",")
	for _, n := range nums {
		i, err := strconv.Atoi(n)
		if err != nil {
			return nil, nil, err
		}
		calls = append(calls, i)
	}

	var (
		b   *board = nil
		row int    = 0
	)
	// Parse bingo boards; assume 5 rows per board
	for _, line := range lines[2:] {
		if b == nil {
			b = newBoard()
		}

		if len(line) > 0 {
			nums := strings.Fields(line)
			for col, num := range nums {
				i, err := strconv.Atoi(num)
				if err != nil {
					return nil, nil, err
				}
				b.addSpace(i, row, col)
			}
			row++
		}

		if row == 5 {
			boards = append(boards, b)
			b = nil
			row = 0
		}
	}

	return calls, boards, nil
}

func part1(calls []int, boards []*board) int {

	for _, call := range calls {
		for _, b := range boards {
			b.addCall(call)

			if b.hasBingo() {
				return b.score * call
			}
		}
	}

	return -1
}

func part2(calls []int, boards []*board) int {
	bingoes := make(map[int]bool)
	for _, call := range calls {
		for i, b := range boards {
			b.addCall(call)

			if b.hasBingo() {
				bingoes[i] = true
			}

			if len(bingoes) == len(boards) {
				return b.score * call
			}
		}
	}

	return -1
}

func main() {
	_, path, _, _ := runtime.Caller(0)
	path = filepath.Join(filepath.Dir(path), "input.txt")

	calls, boards, err := readInput(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(calls, boards))

	for _, b := range boards {
		b.reset()
	}

	fmt.Printf("Part 2: %d\n", part2(calls, boards))
}
