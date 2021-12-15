package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSpace(t *testing.T) {
	assert := assert.New(t)

	b := newBoard()

	spaces := [][]int{
		{1, 0, 0},
		{2, 0, 1},
		{3, 1, 0},
		{4, 1, 1},
	}

	expectedScore := 10

	for _, s := range spaces {
		b.addSpace(s[0], s[1], s[2])
	}

	assert.Equal(4, len(b.nums), "board should have 4 spaces")
	assert.Equal(expectedScore, b.score, "score should be equal to the total of all spaces")
}

func TestAddCall(t *testing.T) {
	assert := assert.New(t)

	b := newBoard()

	spaces := [][]int{
		{1, 0, 0},
		{2, 0, 1},
		{3, 0, 2},
		{4, 0, 3},
		{5, 0, 4},
		{6, 1, 0},
		{7, 1, 1},
		{8, 1, 2},
		{9, 1, 3},
		{10, 1, 4},
		{11, 2, 0},
		{12, 2, 1},
		{13, 2, 2},
		{14, 2, 3},
		{15, 2, 4},
		{16, 3, 0},
		{17, 3, 1},
		{18, 3, 2},
		{19, 3, 3},
		{20, 3, 4},
		{21, 4, 0},
		{22, 4, 1},
		{23, 4, 2},
		{24, 4, 3},
		{25, 4, 4},
	}

	for _, s := range spaces {
		b.addSpace(s[0], s[1], s[2])
	}

	b.addCall(99)

	expectedScore := 325

	for _, p := range b.patterns {
		assert.Zero(p, "all patterns should have zero entries if called number not on board")
	}
	assert.Equal(expectedScore, b.score, "score should be unchanged if called number not on board")

	b.addCall(13)
	assert.Equal(1, b.patterns[2], "row 2 should be marked with a hit")
	assert.Equal(1, b.patterns[7], "column 2 should be marked with a hit")
	assert.Equal(expectedScore-13, b.score, "score should be decreased by called number if found on board")
}

func TestHasBingo(t *testing.T) {
	assert := assert.New(t)

	b := newBoard()

	spaces := [][]int{
		{1, 0, 0},
		{2, 0, 1},
		{3, 0, 2},
		{4, 0, 3},
		{5, 0, 4},
		{6, 1, 0},
		{7, 1, 1},
		{8, 1, 2},
		{9, 1, 3},
		{10, 1, 4},
		{11, 2, 0},
		{12, 2, 1},
		{13, 2, 2},
		{14, 2, 3},
		{15, 2, 4},
		{16, 3, 0},
		{17, 3, 1},
		{18, 3, 2},
		{19, 3, 3},
		{20, 3, 4},
		{21, 4, 0},
		{22, 4, 1},
		{23, 4, 2},
		{24, 4, 3},
		{25, 4, 4},
	}

	for _, s := range spaces {
		b.addSpace(s[0], s[1], s[2])
	}

	b.addCall(1)
	b.addCall(3)
	b.addCall(5)
	b.addCall(7)
	b.addCall(9)
	b.addCall(11)
	b.addCall(13)
	b.addCall(15)
	b.addCall(17)
	b.addCall(19)
	b.addCall(21)
	b.addCall(23)
	b.addCall(25)

	assert.False(b.hasBingo(), "board should not have bingo yet")

	b.addCall(2)
	b.addCall(4)

	assert.True(b.hasBingo(), "board should now have bingo")
}

func TestReset(t *testing.T) {
	assert := assert.New(t)

	b := newBoard()

	spaces := [][]int{
		{1, 0, 0},
		{2, 0, 1},
		{3, 0, 2},
		{4, 0, 3},
		{5, 0, 4},
		{6, 1, 0},
		{7, 1, 1},
		{8, 1, 2},
		{9, 1, 3},
		{10, 1, 4},
		{11, 2, 0},
		{12, 2, 1},
		{13, 2, 2},
		{14, 2, 3},
		{15, 2, 4},
		{16, 3, 0},
		{17, 3, 1},
		{18, 3, 2},
		{19, 3, 3},
		{20, 3, 4},
		{21, 4, 0},
		{22, 4, 1},
		{23, 4, 2},
		{24, 4, 3},
		{25, 4, 4},
	}

	for _, s := range spaces {
		b.addSpace(s[0], s[1], s[2])
	}

	b.addCall(1)
	b.addCall(3)
	b.addCall(5)
	b.addCall(7)
	b.addCall(9)
	b.addCall(11)
	b.addCall(13)
	b.addCall(15)
	b.addCall(17)
	b.addCall(19)
	b.addCall(21)
	b.addCall(23)
	b.addCall(25)

	b.reset()

	for _, p := range b.patterns {
		assert.Zero(p, "all patterns should have zero entries after reset")
	}
	assert.Equal(325, b.score, "score should be reset to original total")
}
