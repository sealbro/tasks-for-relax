package counting_bits

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := 2
	expected := []int{0, 1, 1}

	actual := countBits(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := 5
	expected := []int{0, 1, 1, 2, 1, 2}

	actual := countBits(input)

	assert.EqualMany(t, expected, actual)
}
