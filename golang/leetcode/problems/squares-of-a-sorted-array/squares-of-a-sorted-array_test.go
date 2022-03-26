package squares_of_a_sorted_array

import "testing"

func TestSearchCase1(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	expectedArr := []int{0, 1, 9, 16, 100}

	actualArr := sortedSquares(nums)

	for i, actual := range actualArr {
		expected := expectedArr[i]
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d.", expected, actual)
		}
	}
}
