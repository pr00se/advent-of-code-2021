package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/pr00se/advent-of-code-2021/data"
)

// stack implements a LIFO queue
type stack []rune

// push adds an item to the stack
func (s *stack) push(i rune) {
	*s = append(*s, i)
}

// pop removes and returns an item from the stack. If stack is empty,
// second value returned is false
func (s *stack) pop() (rune, bool) {
	if len(*s) == 0 {
		return '!', false
	}

	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r, true
}

// isOpeningChar returns true if c is a chunk opening character
func isOpeningChar(c rune) bool {
	switch c {
	case '{', '[', '(', '<':
		return true
	}
	return false
}

// getMate returns the corresponding opener/closer to c
func getMate(c rune) rune {
	switch c {
	case '(':
		return ')'
	case ')':
		return '('
	case '{':
		return '}'
	case '}':
		return '{'
	case '[':
		return ']'
	case ']':
		return '['
	case '<':
		return '>'
	case '>':
		return '<'
	default:
		return '!'
	}
}

// corrumptedScore returns the score for c in a corrupted line
func corruptedScore(c rune) int {
	switch c {
	case ')':
		return 3
	case '}':
		return 1197
	case ']':
		return 57
	case '>':
		return 25137
	default:
		return -1
	}
}

// incompleteScore returns the score for c in an incomplete line
func incompleteScore(c rune) int {
	switch c {
	case ')':
		return 1
	case '}':
		return 3
	case ']':
		return 2
	case '>':
		return 4
	default:
		return -1
	}
}

// detectCorruptChunks returns the corruption score for the line if it is
// corrupted, and 0 if it is not
func detectCorruptChunks(line string) int {
	var openers stack

	for _, c := range strings.TrimSpace(line) {
		// assume that we'll never get extra closing characters and
		// run out of openers...
		if isOpeningChar(c) {
			openers.push(c)
		} else if r, ok := openers.pop(); ok && r != getMate(c) {
			return corruptedScore(c)
		}
	}

	return 0
}

// fixIncompleteChunks returns the score for completing the incomplete line
func fixIncompleteChunks(line string) int {
	var openers stack

	for _, c := range strings.TrimSpace(line) {
		// assume we're always working with uncorrupted lines so
		// don't bother checking whether pairs match
		if isOpeningChar(c) {
			openers.push(c)
		} else {
			openers.pop()
		}
	}

	var score int
	for r, ok := openers.pop(); ok; r, ok = openers.pop() {
		score = (score * 5) + incompleteScore(getMate(r))
	}

	return score
}

func part1(lines []string) int {
	var score int

	for _, l := range lines {
		score += detectCorruptChunks(l)
	}

	return score
}

func part2(lines []string) int {
	var (
		scores     []int
		incomplete []string
	)

	for _, l := range lines {
		if s := detectCorruptChunks(l); s == 0 {
			incomplete = append(incomplete, l)
		}
	}

	for _, l := range incomplete {
		scores = append(scores, fixIncompleteChunks(l))
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func main() {
	input, err := data.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 1: %d\n", part2(lines))
}
