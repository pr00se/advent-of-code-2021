package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommaSeparatedInts(t *testing.T) {
	assert := assert.New(t)

	input := `     16,	1,2,   0,	4   ,2,74,1,2, 14 	`

	expected := []int{16, 1, 2, 0, 4, 2, 74, 1, 2, 14}

	output, err := ParseCommaSeparatedInts(input)

	assert.Nil(err, "error should be nil")
	assert.Equalf(expected, output, "numbers should be parsed correctly")
}
