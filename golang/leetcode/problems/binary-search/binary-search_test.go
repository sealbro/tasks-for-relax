package binary_search

import "testing"

func TestSearchCase1(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4

	actual := search(nums, target)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}

func TestSearchCase2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 9, 12}
	target := 9
	expected := 5

	actual := search(nums, target)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}

func TestSearchCase3(t *testing.T) {
	nums := []int{5}
	target := 5
	expected := 0

	actual := search(nums, target)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}

func TestSearchCase4(t *testing.T) {
	nums := []int{2, 5}
	target := 5
	expected := 1

	actual := search(nums, target)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}
