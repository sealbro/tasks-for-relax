package the_k_weakest_rows_in_a_matrix

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	target := 3
	expected := []int{2, 0, 3}

	actual := kWeakestRows(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	target := 2
	expected := []int{0, 2}

	actual := kWeakestRows(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := [][]int{
		{1, 0},
		{1, 1},
		{0, 0},
	}
	target := 1
	expected := []int{2}

	actual := kWeakestRows(input, target)

	assert.EqualMany(t, expected, actual)
}
