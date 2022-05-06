package sort_integers_by_the_number_of_1_bits

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	expected := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}

	actual := sortByBits(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	expected := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	actual := sortByBits(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{8, 4, 3}
	expected := []int{4, 8, 3}

	actual := sortByBits(input)

	assert.EqualMany(t, expected, actual)
}
