package move_zeroes

import "testing"

func TestSearchCase1(t *testing.T) {
	actual := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}

	moveZeroes(actual)

	for i, actualItem := range actual {
		expectedItem := expected[i]
		if actualItem != expectedItem {
			t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
			break
		}
	}
}

func TestSearchCase2(t *testing.T) {
	actual := []int{0, -1, 0, 1, 0, 3, 12}
	expected := []int{-1, 1, 3, 12, 0, 0, 0}

	moveZeroes(actual)

	for i, actualItem := range actual {
		expectedItem := expected[i]
		if actualItem != expectedItem {
			t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
			break
		}
	}
}

func TestSearchCase3(t *testing.T) {
	actual := []int{0}
	expected := []int{0}

	moveZeroes(actual)

	for i, actualItem := range actual {
		expectedItem := expected[i]
		if actualItem != expectedItem {
			t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
			break
		}
	}
}
