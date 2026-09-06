package remove_duplicates_from_sorted_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 1, 2}
	expected := 2

	k := removeDuplicates(input)

	assert.Equal(t, expected, k)
}

func TestCase2(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := 5

	k := removeDuplicates(input)

	assert.Equal(t, expected, k)
}
