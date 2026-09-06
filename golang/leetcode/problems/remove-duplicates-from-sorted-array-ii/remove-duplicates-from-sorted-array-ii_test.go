package remove_duplicates_from_sorted_array_ii

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 1, 1, 2, 2, 3}
	expectedNums := []int{1, 1, 2, 2, 3}
	expected := 5

	k := removeDuplicates(input)

	assert.Equal(t, expected, k)
	for i := range k {
		assert.Equal(t, expectedNums[i], input[i])
	}
}

func TestCase2(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	expectedNums := []int{0, 0, 1, 1, 2, 3, 3}
	expected := 7

	k := removeDuplicates(input)

	assert.Equal(t, expected, k)
	for i := range k {
		assert.Equal(t, expectedNums[i], input[i])
	}
}
