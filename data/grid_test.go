package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePoint(t *testing.T) {
	assert := assert.New(t)

	input := `     16,	-1 	`

	expected := Point{X: 16, Y: -1}

	output, err := ParsePoint(input)

	assert.Nil(err, "error should be nil")
	assert.Equalf(expected, output, "point should be parsed correctly")
}

func TestParseGrid(t *testing.T) {
	assert := assert.New(t)

	input := `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`

	output, err := ParseGrid(input)

	assert.Nil(err, "error should be nil")
	assert.Equal(50, len(output), "should have parsed 50 points")
}
