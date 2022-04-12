package rotting_oranges

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	matrix := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	expected := 4

	actual := orangesRotting(matrix)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	matrix := [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}

	expected := -1

	actual := orangesRotting(matrix)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	matrix := [][]int{
		{0, 2},
	}

	expected := 0

	actual := orangesRotting(matrix)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	matrix := [][]int{
		{2, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 0, 1, 1},
		{2, 0, 1, 2},
	}

	expected := 2

	actual := orangesRotting(matrix)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	matrix := [][]int{
		{2, 1, 1},
		{1, 1, 1},
		{0, 1, 2},
	}
	expected := 2

	actual := orangesRotting(matrix)

	assert.Equal(t, expected, actual)
}
