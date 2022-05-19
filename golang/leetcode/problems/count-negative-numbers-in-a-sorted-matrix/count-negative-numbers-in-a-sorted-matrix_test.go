package count_negative_numbers_in_a_sorted_matrix

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	expected := 8

	actual := countNegatives(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := [][]int{
		{3, 2},
		{1, 0},
	}
	expected := 0

	actual := countNegatives(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := [][]int{
		{0},
	}
	expected := 0

	actual := countNegatives(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := [][]int{
		{-4, -2, -1},
		{-4, -2, -1},
		{-4, -2, -1},
	}
	expected := 9

	actual := countNegatives(nums)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums := [][]int{
		{4, -2, -1},
		{4, -2, -1},
		{4, -2, -1},
	}
	expected := 6

	actual := countNegatives(nums)

	assert.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	nums := [][]int{
		{4, 2, 1},
		{4, 2, -1},
		{4, -2, -1},
	}
	expected := 3

	actual := countNegatives(nums)

	assert.Equal(t, expected, actual)
}
