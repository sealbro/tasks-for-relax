package search_a_2d_matrix

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	nums := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 13

	actual := searchMatrix(nums, target)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	nums := [][]int{
		{1, 3, 5, 7},
	}
	target := 5

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase4(t *testing.T) {
	nums := [][]int{
		{1, 3, 5, 7},
		{10, 11, 13, 20},
		{23, 30, 34, 60},
	}
	target := 13

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase5(t *testing.T) {
	nums := [][]int{
		{1, 3, 5, 7},
	}
	target := 8

	actual := searchMatrix(nums, target)

	assert.False(t, actual)
}

func TestCase6(t *testing.T) {
	nums := [][]int{
		{1},
	}
	target := 8

	actual := searchMatrix(nums, target)

	assert.False(t, actual)
}

func TestCase7(t *testing.T) {
	nums := [][]int{
		{8},
	}
	target := 8

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase8(t *testing.T) {
	nums := [][]int{
		{1},
	}
	target := 1

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase9(t *testing.T) {
	nums := [][]int{
		{1, 3},
	}
	target := 3

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase10(t *testing.T) {
	nums := [][]int{
		{1, 3},
		{5, 7},
	}
	target := 7

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase11(t *testing.T) {
	nums := [][]int{
		{1},
		{3},
	}
	target := 3

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase12(t *testing.T) {
	nums := [][]int{
		{1, 3, 5},
	}
	target := 5

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase13(t *testing.T) {
	nums := [][]int{
		{1, 3, 5},
		{11, 13, 15},
	}
	target := 15

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}

func TestCase14(t *testing.T) {
	nums := [][]int{
		{-9, -7, -7, -5, -3},
		{-1, 0, 1, 3, 3},
		{5, 7, 9, 11, 12},
		{13, 14, 15, 16, 18},
		{19, 21, 22, 22, 22},
	}
	target := -4

	actual := searchMatrix(nums, target)

	assert.False(t, actual)
}

func TestCase15(t *testing.T) {
	nums := [][]int{
		{-9, -7, -7, -5, -3},
		{-1, 0, 1, 3, 3},
		{5, 7, 9, 11, 12},
		{13, 14, 15, 16, 18},
		{19, 21, 22, 22, 22},
	}
	target := 22

	actual := searchMatrix(nums, target)

	assert.True(t, actual)
}
