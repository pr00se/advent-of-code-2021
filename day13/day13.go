package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

type fold struct {
	axis rune
	val  int
}

// parseInput parses the input string and returns the grid and folds
func parseInput(input string) (data.Grid, []fold, error) {
	var (
		grid  = data.Grid{}
		folds []fold
	)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		if len(line) > 0 {
			p, err := data.ParsePoint(line)
			if err != nil {
				f := fold{}
				_, err := fmt.Sscanf(strings.TrimSpace(line), "fold along %c=%d", &f.axis, &f.val)
				if err != nil {
					return nil, nil, err
				}
				folds = append(folds, f)
			} else {
				grid[p] = 0
			}
		}
	}

	return grid, folds, nil
}

// foldGrid folds the grid along fold, updating the locations of affected points
func foldGrid(grid data.Grid, f fold) {
	var toDelete []data.Point

	for p := range grid {
		if f.axis == 'x' {
			if p.X > f.val {
				ref := data.Point{X: f.val - (p.X - f.val), Y: p.Y}
				grid[ref] = 0
				toDelete = append(toDelete, p)
			}
		} else {
			if p.Y > f.val {
				ref := data.Point{X: p.X, Y: f.val - (p.Y - f.val)}
				grid[ref] = 0
				toDelete = append(toDelete, p)
			}
		}
	}

	for _, p := range toDelete {
		delete(grid, p)
	}
}

func part1(grid data.Grid, folds []fold) int {
	foldGrid(grid, folds[0])
	return len(grid)
}

func part2(grid data.Grid, folds []fold) string {
	for _, f := range folds {
		foldGrid(grid, f)
	}
	return grid.String()
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	grid, folds, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(grid, folds))
	// part 1 executed the first fold, so just execute the rest
	fmt.Printf("Part 2: \n%s\n", part2(grid, folds[1:]))
}
