package pascals_triangle

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 5
	expected := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}

	actual := generate(target)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := 1
	expected := [][]int{{1}}

	actual := generate(target)

	assert.EqualMultiDimensional(t, expected, actual)
}
