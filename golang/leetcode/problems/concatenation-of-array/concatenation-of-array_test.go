package concatenation_of_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 2, 1}
	expected := []int{1, 2, 1, 1, 2, 1}

	actual := getConcatenation(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{1, 3, 2, 1}
	expected := []int{1, 3, 2, 1, 1, 3, 2, 1}

	actual := getConcatenation(input)

	assert.EqualMany(t, expected, actual)
}
