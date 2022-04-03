package max_area_of_island

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	expected := 6

	actual := maxAreaOfIsland(grid)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	expected := 0

	actual := maxAreaOfIsland(grid)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}

	expected := 4

	actual := maxAreaOfIsland(grid)

	assert.Equal(t, expected, actual)
}
