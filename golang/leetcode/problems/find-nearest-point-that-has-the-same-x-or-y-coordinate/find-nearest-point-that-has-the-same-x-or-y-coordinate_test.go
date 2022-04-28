package find_nearest_point_that_has_the_same_x_or_y_coordinate

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	x, y := 3, 4
	points := [][]int{
		{1, 2},
		{3, 1},
		{2, 4},
		{2, 3},
		{4, 4},
	}

	expected := 2

	actual := nearestValidPoint(x, y, points)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	x, y := 3, 4
	points := [][]int{
		{3, 4},
	}

	expected := 0

	actual := nearestValidPoint(x, y, points)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	x, y := 3, 4
	points := [][]int{
		{2, 3},
	}

	expected := -1

	actual := nearestValidPoint(x, y, points)

	assert.Equal(t, expected, actual)
}
