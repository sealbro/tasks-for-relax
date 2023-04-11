package build_array_from_permutation

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{0, 2, 1, 5, 3, 4}
	expected := []int{0, 1, 2, 4, 5, 3}

	actual := buildArray(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{5, 0, 1, 2, 3, 4}
	expected := []int{4, 5, 0, 1, 2, 3}

	actual := buildArray(input)

	assert.EqualMany(t, expected, actual)
}
