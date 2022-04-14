package triangle

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	expected := 11

	actual := minimumTotal(input)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := [][]int{
		{-10},
	}
	expected := -10

	actual := minimumTotal(input)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := [][]int{
		{2},
		{3, 4},
		{5, 5, 7},
		{-20, 1, 8, 3},
	}
	expected := -10

	actual := minimumTotal(input)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := [][]int{
		{-1},
		{2, 3},
		{1, -1, -3},
	}
	expected := -1

	actual := minimumTotal(input)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := [][]int{
		{-1},
		{3, 2},
		{-3, 1, -1},
	}
	expected := -1

	actual := minimumTotal(input)

	assert.Equal(t, expected, actual)
}
