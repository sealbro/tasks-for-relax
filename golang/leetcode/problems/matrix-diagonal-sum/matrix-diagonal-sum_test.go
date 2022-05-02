package matrix_diagonal_sum

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := 25

	actual := diagonalSum(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	expected := 8

	actual := diagonalSum(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := [][]int{
		{5},
	}

	expected := 5

	actual := diagonalSum(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}

	expected := 9

	actual := diagonalSum(nums)

	assert.Equal(t, expected, actual)
}
