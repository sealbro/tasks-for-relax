package check_if_it_is_a_straight_line

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7},
	}

	actual := checkStraightLine(nums)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	nums := [][]int{
		{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7},
	}

	actual := checkStraightLine(nums)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	nums := [][]int{
		{1, 1}, {2, 2},
	}

	actual := checkStraightLine(nums)

	assert.True(t, actual)
}

func TestCase4(t *testing.T) {
	nums := [][]int{
		{1, 1}, {2, 2}, {6, 6}, {7, 7}, {3, 4}, {4, 5},
	}

	actual := checkStraightLine(nums)

	assert.False(t, actual)
}

func TestCase5(t *testing.T) {
	nums := [][]int{
		{0, 0}, {0, 1}, {0, -1},
	}

	actual := checkStraightLine(nums)

	assert.True(t, actual)
}

func TestCase6(t *testing.T) {
	nums := [][]int{
		{2, 1}, {4, 2}, {6, 3},
	}

	actual := checkStraightLine(nums)

	assert.True(t, actual)
}
