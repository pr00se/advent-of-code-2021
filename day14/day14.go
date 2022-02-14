package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

type insertionRules map[string]rune

// parseInput parses the input string and returns the template and insertion rules
func parseInput(input string) (string, insertionRules, error) {
	var (
		template string
		rules    = insertionRules{}
	)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	template = strings.TrimSpace(lines[0])

	for _, line := range lines[1:] {
		if len(line) > 0 {
			tokens := strings.Fields(line)

			if len(tokens) != 3 {
				return "", nil, fmt.Errorf("invalid input: %s", line)
			}

			rules[tokens[0]] = rune(tokens[2][0])
		}
	}

	return template, rules, nil
}

func insertPairs(pairs map[string]int, rules insertionRules) map[string]int {
	output := map[string]int{}

	for p, count := range pairs {
		element := rules[p]
		output[string(p[0])+string(element)] += count
		output[string(element)+string(p[1])] += count
	}

	return output
}

func part1(template string, rules insertionRules) int {
	pairs := map[string]int{}

	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for i := 0; i < 10; i++ {
		pairs = insertPairs(pairs, rules)
	}

	counts := map[byte]int{}

	for p, count := range pairs {
		counts[p[0]] += count
	}

	counts[template[len(template)-1]]++

	var max, min int
	for _, count := range counts {
		if max == 0 || count > max {
			max = count
		}
		if min == 0 || count < min {
			min = count
		}
	}

	return max - min
}

func part2(template string, rules insertionRules) int {
	pairs := map[string]int{}

	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for i := 0; i < 40; i++ {
		pairs = insertPairs(pairs, rules)
	}

	counts := map[byte]int{}

	for p, count := range pairs {
		counts[p[0]] += count
	}

	counts[template[len(template)-1]]++

	var max, min int
	for _, count := range counts {
		if max == 0 || count > max {
			max = count
		}
		if min == 0 || count < min {
			min = count
		}
	}

	return max - min
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	template, rules, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(template, rules))
	fmt.Printf("Part 2: %d\n", part2(template, rules))
}
