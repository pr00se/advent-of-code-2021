package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

type command struct {
	op  string
	val int
}

func parseInput(input string) ([]command, error) {
	var cmds []command

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, l := range lines {
		cmd := command{}

		_, err := fmt.Sscanf(l, "%s %d", &cmd.op, &cmd.val)
		if err != nil {
			return nil, err
		}

		cmds = append(cmds, cmd)
	}

	return cmds, nil
}

func part1(cmds []command) int {
	var hpos, depth int

	for _, c := range cmds {
		switch c.op {
		case "forward":
			hpos += c.val
		case "down":
			depth += c.val
		case "up":
			depth -= c.val
		}
	}

	return hpos * depth
}

func part2(cmds []command) int {
	var hpos, depth, aim int

	for _, c := range cmds {
		switch c.op {
		case "forward":
			hpos += c.val
			depth += (aim * c.val)
		case "down":
			aim += c.val
		case "up":
			aim -= c.val
		}
	}

	return hpos * depth
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	cmds, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(cmds))
	fmt.Printf("Part 2: %d\n", part2(cmds))
}
