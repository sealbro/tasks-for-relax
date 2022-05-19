package search_in_rotated_sorted_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	expected := 4

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	expected := -1

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{1}
	target := 0
	expected := -1

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{1}
	target := 0
	expected := -1

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 5
	expected := 1

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 6
	expected := 2

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase8(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 2
	expected := 6

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase9(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 1
	expected := 5

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase10(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 1, 2, 3}
	target := 8
	expected := 4

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase11(t *testing.T) {
	nums := []int{5, 1, 2, 3, 4}
	target := 1
	expected := 1

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase12(t *testing.T) {
	nums := []int{1, 2, 3, 4, 0}
	target := 4
	expected := 3

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}
