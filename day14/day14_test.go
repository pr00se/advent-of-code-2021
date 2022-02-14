package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	input := `NNCB

	CH -> B
	HH -> N
	CB -> H
	NH -> C
	HB -> C
	HC -> B
	HN -> C
	NN -> C
	BH -> H
	NC -> B
	NB -> B
	BN -> B
	BB -> N
	BC -> B
	CC -> N
	CN -> C`

	template, rules, err := parseInput(input)

	assert.Nil(err, "error should be nil")
	assert.Equalf(4, len(template), "should have parsed template correctly")
	assert.Equalf(16, len(rules), "should have parsed 16 insertion rules")
}

func TestPart1(t *testing.T) {
	input := `NNCB

	CH -> B
	HH -> N
	CB -> H
	NH -> C
	HB -> C
	HC -> B
	HN -> C
	NN -> C
	BH -> H
	NC -> B
	NB -> B
	BN -> B
	BB -> N
	BC -> B
	CC -> N
	CN -> C`

	expected := 1588

	template, rules, _ := parseInput(input)

	output := part1(template, rules)

	assert.Equal(t, expected, output)
}

func TestPart2(t *testing.T) {
	input := `NNCB

	CH -> B
	HH -> N
	CB -> H
	NH -> C
	HB -> C
	HC -> B
	HN -> C
	NN -> C
	BH -> H
	NC -> B
	NB -> B
	BN -> B
	BB -> N
	BC -> B
	CC -> N
	CN -> C`

	expected := 2188189693529

	template, rules, _ := parseInput(input)

	output := part2(template, rules)

	assert.Equal(t, expected, output)
}
