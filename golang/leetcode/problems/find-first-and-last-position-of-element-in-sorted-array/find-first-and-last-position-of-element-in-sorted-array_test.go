package find_first_and_last_position_of_element_in_sorted_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{5, 7, 7, 8, 8, 10}
	target := 6
	expected := []int{-1, -1}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{}
	target := 0
	expected := []int{-1, -1}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := []int{6, 7, 7, 8, 8, 10}
	target := 6
	expected := []int{0, 0}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := []int{6, 7, 7, 8, 8, 10}
	target := 10
	expected := []int{5, 5}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase6(t *testing.T) {
	input := []int{6, 7, 7, 8, 8, 10}
	target := 2
	expected := []int{-1, -1}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase7(t *testing.T) {
	input := []int{7}
	target := 7
	expected := []int{0, 0}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase8(t *testing.T) {
	input := []int{10}
	target := 20
	expected := []int{-1, -1}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase9(t *testing.T) {
	input := []int{10}
	target := 6
	expected := []int{-1, -1}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase10(t *testing.T) {
	input := []int{6, 7, 7, 8, 8, 10}
	target := 6
	expected := []int{0, 0}

	actual := searchRange(input, target)

	assert.EqualMany(t, expected, actual)
}
