package main

// A space represents the coordinates of a space on a bingo board
type space struct {
	row int
	col int
}

// A board represents a 5x5 bingo board
type board struct {
	nums     map[int]space // mapping of numbers to their spaces on this board
	patterns []int         // counts of called numbers in winning patterns (rows, columns)
	score    int           // sum of all numbers not yet called
}

// newBoard returns an empty board
func newBoard() *board {
	b := &board{}

	b.nums = make(map[int]space)
	b.patterns = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	b.score = 0

	return b
}

// addSpace adds a new space to the board. Assumes each space has a unique value.
func (b *board) addSpace(val, row, col int) {
	s := space{
		row: row,
		col: col,
	}
	b.nums[val] = s
	b.score += val
}

// addCall records a new called number on the board, if it is present. Repeated calls of the
// same number will be recorded, leading to incorrect scoring.
func (b *board) addCall(val int) {
	if s, ok := b.nums[val]; ok {
		// mark the row
		b.patterns[s.row]++

		// mark the column
		b.patterns[5+s.col]++

		b.score -= val
	}
}

// hasBingo checks if the board has bingo
func (b *board) hasBingo() bool {
	for _, p := range b.patterns {
		if p == 5 {
			return true
		}
	}
	return false
}

// reset resets the board
func (b *board) reset() {
	b.patterns = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	score := 0
	for k := range b.nums {
		score += k
	}
	b.score = score
}
