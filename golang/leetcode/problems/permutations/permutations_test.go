package permutations

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	actual := permute(input)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{0, 1}
	expected := [][]int{
		{0, 1},
		{1, 0},
	}

	actual := permute(input)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{1}
	expected := [][]int{
		{1},
	}

	actual := permute(input)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := []int{0, -1, 1}
	expected := [][]int{
		{0, -1, 1},
		{0, 1, -1},
		{-1, 0, 1},
		{-1, 1, 0},
		{1, 0, -1},
		{1, -1, 0},
	}

	actual := permute(input)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := []int{0, -1, 1}
	expected := [][]int{
		{-1, 0, 1},
		{-1, 1, 0},
		{0, -1, 1},
		{0, 1, -1},
		{1, -1, 0},
		{1, 0, -1},
	}

	actual := permuteLexicographic(input)

	assert.EqualMultiDimensional(t, expected, actual)
}
