package main

import (
	"fmt"
	"log"

	"github.com/pr00se/advent-of-code-2021/data"
)

func part1(cave data.Grid) int {
	risks := data.Grid{}
	start := data.Point{X: 0, Y: 0}

	risks[start] = 0

	var x, y int
	for visiting, toVisit := []data.Point{start}, []data.Point{}; len(visiting) > 0; visiting, toVisit = toVisit, nil {
		for _, v := range visiting {
			risk, ok := risks[v]
			if !ok {
				continue
			}

			if x == 0 || v.X > x {
				x = v.X
			}
			if y == 0 || v.Y > y {
				y = v.Y
			}

			for _, n := range v.Adjacent() {
				if nr, ok := cave[n]; ok {
					if risks[n] == 0 || risk+nr < risks[n] {
						risks[n] = risk + nr
						toVisit = append(toVisit, n)
					}
				}
			}
		}
	}

	drawPath(risks)

	return risks[data.Point{X: x, Y: y}]
}

func constructMap(grid data.Grid) data.Grid {
	output := data.Grid{}

	var maxX, maxY int
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			for p, v := range grid {
				if maxX == 0 || p.X > maxX {
					maxX = p.X
				}

				if maxY == 0 || p.Y > maxY {
					maxY = p.Y
				}

				x := (col * (maxX + 1)) + p.X
				y := (row * (maxY + 1)) + p.Y

				output[data.Point{X: x, Y: y}] = ((v - 1 + row + col) % 9) + 1
			}
		}
	}

	return output
}

func drawPath(risks data.Grid) {
	path := data.Grid{}

	var maxX, maxY int
	for p, _ := range risks {
		if maxX == 0 || p.X > maxX {
			maxX = p.X
		}

		if maxY == 0 || p.Y > maxY {
			maxY = p.Y
		}
	}

	start := data.Point{X: maxX, Y: maxY}

	for curr := start; ; {
		for _, n := range curr.Adjacent() {
			if n.X == 0 && n.Y == 0 {
				path[data.Point{X: 0, Y: 0}] = 0
				fmt.Println(path)
				return
			}

			if r, ok := risks[n]; ok && r < risks[curr] {
				curr = n
			}
		}
		path[curr] = 0
	}
}

func part2(cave data.Grid) int {
	out := constructMap(cave)

	return part1(out)
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	grid, err := data.ParseGrid(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(grid))
	fmt.Printf("Part 2: %d\n", part2(grid))
}
