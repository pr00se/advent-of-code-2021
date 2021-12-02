package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type command struct {
	op  string
	val int
}

func readCommands(path string) ([]command, error) {
	var cmds []command

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	for _, l := range lines {
		if len(l) > 0 {
			cmd := command{}

			_, err := fmt.Sscanf(l, "%s %d", &cmd.op, &cmd.val)
			if err != nil {
				return nil, err
			}

			cmds = append(cmds, cmd)
		}
	}

	return cmds, nil
}

func part1(cmds []command) int {
	hpos, depth := 0, 0

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
	hpos, depth, aim := 0, 0, 0

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
	_, path, _, _ := runtime.Caller(0)
	path = filepath.Join(filepath.Dir(path), "input.txt")

	cmds, err := readCommands(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(cmds))
	fmt.Printf("Part 2: %d\n", part2(cmds))
}
