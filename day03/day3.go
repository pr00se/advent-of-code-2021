package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

func part1(lines []string) int {
	counts := make([]int, len(lines[0]))

	for _, l := range lines {
		for i, c := range l {
			if c == '1' {
				counts[i]++
			}
		}
	}

	var g, e int

	for _, c := range counts {
		g = g << 1
		e = e << 1
		if c >= len(lines)/2 {
			// more 1s than 0s; gamma has a 1, epsilon a 0
			g += 1
		} else {
			// more 0s than 1s; gamma has a 0, epsilon a 1
			e += 1
		}
	}

	return g * e
}

func part2(lines []string) int {
	var candidates []string

	lists := map[byte][]string{
		'0': {},
		'1': {},
	}

	candidates = lines
	for i := 0; i < len(lines[0]); i++ {
		if len(candidates) == 0 {
			log.Fatal("No valid candidates for O2")
		}

		for x := range lists {
			lists[x] = nil
		}

		for _, c := range candidates {
			lists[c[i]] = append(lists[c[i]], c)
		}

		if len(lists['1']) >= len(lists['0']) {
			// more 1s than 0s; keep list of 1s
			candidates = lists['1']
		} else {
			// more 0s than 1s; keep list of 0s
			candidates = lists['0']
		}

		if len(candidates) == 1 {
			break
		}
	}

	if len(candidates) != 1 {
		log.Fatalf("Expected a single candidate for O2, got %d", len(candidates))
	}

	i, err := strconv.ParseInt(candidates[0], 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	o2 := int(i)

	candidates = lines
	for i := 0; i < len(lines[0]); i++ {
		if len(candidates) < 1 {
			log.Fatal("No valid candidates for CO2")
		}

		for x := range lists {
			lists[x] = nil
		}

		for _, c := range candidates {
			lists[c[i]] = append(lists[c[i]], c)
		}

		if len(lists['1']) >= len(lists['0']) {
			// more 1s than 0s; keep list of 0s
			candidates = lists['0']
		} else {
			// more 0s than 1s; keep list of 1s
			candidates = lists['1']
		}

		if len(candidates) == 1 {
			break
		}
	}

	if len(candidates) != 1 {
		log.Fatalf("Expected a single candidate for CO2, got %d", len(candidates))
	}

	i, err = strconv.ParseInt(candidates[0], 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	co2 := int(i)

	return o2 * co2
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
