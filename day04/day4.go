package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

// parseInput parses the input string and returns the bingo calls and boards
func parseInput(input string) ([]int, []*board, error) {
	var (
		calls  []int
		boards []*board
	)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Parse list of bingo calls
	calls, err := data.ParseCommaSeparatedInts(lines[0])
	if err != nil {
		return nil, nil, err
	}

	var b *board
	var row int
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
	bingoes := map[int]bool{}
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
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	calls, boards, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(calls, boards))

	for _, b := range boards {
		b.reset()
	}

	fmt.Printf("Part 2: %d\n", part2(calls, boards))
}
