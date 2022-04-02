package flood_fill

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	image := [][]int{{0, 0, 0}, {0, 0, 0}}
	sr := 0
	sc := 0
	newColor := 2

	expected := [][]int{{2, 2, 2}, {2, 2, 2}}

	actual := floodFill(image, sr, sc, newColor)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase2(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2

	expected := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}

	actual := floodFill(image, sr, sc, newColor)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase3(t *testing.T) {
	image := [][]int{{0, 0, 0}, {0, 1, 1}}
	sr := 1
	sc := 1
	newColor := 1

	expected := [][]int{{0, 0, 0}, {0, 1, 1}}

	actual := floodFill(image, sr, sc, newColor)

	assert.EqualMultiDimensional(t, expected, actual)
}
