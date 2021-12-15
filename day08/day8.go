package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

type display struct {
	patterns []string
	codes    []string
}

// parseInput parses the input string and returns the displays
func parseInput(input string) ([]display, error) {
	var displays []display

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		d := display{}

		halves := strings.Split(strings.TrimSpace(line), "|")

		d.patterns = strings.Fields(halves[0])
		d.codes = strings.Fields(halves[1])

		displays = append(displays, d)
	}

	return displays, nil
}

// decode correlates patterns with digits and returns the displayed code
func decode(d display) int {
	candidates := map[int][]int{}
	encoded := map[int]int{}
	decoded := map[int]int{}

	for _, p := range d.patterns {
		e := encodePattern(p)

		switch len(p) {
		case 2:
			// pattern of length 2 is 1
			encoded[1] = e
			decoded[e] = 1
		case 3:
			// pattern of length 3 is 7
			encoded[7] = e
			decoded[e] = 7
		case 4:
			// pattern of length 4 is 4
			encoded[4] = e
			decoded[e] = 4
		case 5:
			candidates[5] = append(candidates[5], e)
		case 6:
			candidates[6] = append(candidates[6], e)
		case 7:
			// pattern of length 7 is 8
			encoded[8] = e
			decoded[e] = 8
		}
	}

	for _, c := range candidates[6] {
		if (c & encoded[4]) == encoded[4] {
			// pattern of length 6 that contains all of pattern 4 is 9
			encoded[9] = c
			decoded[c] = 9
		} else if (c & encoded[1]) != encoded[1] {
			// pattern of length 6 that does not contain all of pattern 1 is 6
			encoded[6] = c
			decoded[c] = 6
		} else {
			// the remaining length 6 pattern is 0
			encoded[0] = c
			decoded[c] = 0
		}
	}

	for _, c := range candidates[5] {
		if (c & encoded[1]) == encoded[1] {
			// pattern of length 5 that contains all of pattern 1 is 3
			encoded[3] = c
			decoded[c] = 3
		} else if (c & encoded[6]) == c {
			// pattern of length 5 that is contained within pattern 6 is 5
			encoded[5] = c
			decoded[c] = 5
		} else {
			// the remaining length 5 pattern is 2
			encoded[2] = c
			decoded[c] = 2
		}
	}

	var output int
	for _, c := range d.codes {
		output = (output * 10) + decoded[encodePattern(c)]
	}

	return output
}

// encodePattern encodes a signal pattern into a bitmask
func encodePattern(s string) int {
	var out int

	for _, c := range s {
		switch c {
		case 'a':
			out |= 0b1000000
		case 'b':
			out |= 0b0100000
		case 'c':
			out |= 0b0010000
		case 'd':
			out |= 0b0001000
		case 'e':
			out |= 0b0000100
		case 'f':
			out |= 0b0000010
		case 'g':
			out |= 0b0000001
		}
	}

	return out
}

func part1(displays []display) int {
	var count int

	for _, d := range displays {
		for _, c := range d.codes {
			// only non-unique lengths are 5 and 6
			if len(c) != 5 && len(c) != 6 {
				count++
			}
		}
	}

	return count
}

func part2(displays []display) int {
	var total int

	for _, d := range displays {
		total += decode(d)
	}

	return total
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	displays, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(displays))
	fmt.Printf("Part 2: %d\n", part2(displays))
}
