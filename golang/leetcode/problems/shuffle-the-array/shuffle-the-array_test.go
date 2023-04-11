package shuffle_the_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{2, 5, 1, 3, 4, 7}
	n := 3
	expected := []int{2, 3, 5, 4, 1, 7}

	actual := shuffle(input, n)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n := 4
	expected := []int{1, 4, 2, 3, 3, 2, 4, 1}

	actual := shuffle(input, n)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{1, 1, 2, 2}
	n := 2
	expected := []int{1, 2, 1, 2}

	actual := shuffle(input, n)

	assert.EqualMany(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := []int{1, 2, 3, 1, 2, 3}
	n := 3
	expected := []int{1, 1, 2, 2, 3, 3}

	actual := shuffle(input, n)

	assert.EqualMany(t, expected, actual)
}
