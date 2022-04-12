package combinations

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	n := 4
	k := 2

	expected := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{1, 4},
		{2, 4},
		{3, 4},
	}

	actual := combine(n, k)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase2(t *testing.T) {
	n := 1
	k := 1

	expected := [][]int{
		{1},
	}

	actual := combine(n, k)

	assert.EqualMultiDimensional(t, expected, actual)
}
