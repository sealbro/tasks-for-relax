package search_in_a_sorted_array_of_unknown_size

import (
	"testing"
)

func TestSearchCase1(t *testing.T) {
	nums := ArrayReader{
		bucket: []int{-1, 0, 3, 5, 9, 12},
	}
	target := 9
	expected := 4

	actual := search(nums, target)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}
