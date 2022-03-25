package search_insert_position

import "testing"

func TestSearchCase1(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	expected := 2

	actual := searchInsert(nums, target)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}

func TestSearchCase2(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	expected := 1

	actual := searchInsert(nums, target)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}

func TestSearchCase3(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	expected := 4

	actual := searchInsert(nums, target)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}
