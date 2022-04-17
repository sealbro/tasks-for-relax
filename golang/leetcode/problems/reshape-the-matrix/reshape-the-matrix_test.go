package reshape_the_matrix

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	r := 1
	c := 4
	expected := [][]int{{1, 2, 3, 4}}

	actual := matrixReshape(matrix, r, c)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase2(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	r := 2
	c := 4
	expected := [][]int{
		{1, 2},
		{3, 4},
	}

	actual := matrixReshape(matrix, r, c)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase3(t *testing.T) {
	matrix := [][]int{
		{1, 2},
	}
	r := 1
	c := 1
	expected := [][]int{
		{1, 2},
	}

	actual := matrixReshape(matrix, r, c)

	assert.EqualMultiDimensional(t, expected, actual)
}
