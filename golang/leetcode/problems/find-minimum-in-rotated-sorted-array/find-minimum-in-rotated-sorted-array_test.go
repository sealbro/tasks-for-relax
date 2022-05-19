package find_minimum_in_rotated_sorted_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	expected := 1

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	expected := 0

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{11, 13, 15, 17}
	expected := 11

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{12, 13, 15, 11}
	expected := 11

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums := []int{4, 1, 2, 3}
	expected := 1

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	nums := []int{5, 1, 2, 3, 4}
	expected := 1

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}

func TestCase7(t *testing.T) {
	nums := []int{5, 6, 1, 2, 3, 4}
	expected := 1

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}

func TestCase8(t *testing.T) {
	nums := []int{2, 3, 4, 1}
	expected := 1

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}

func TestCase9(t *testing.T) {
	nums := []int{2, 3, 4, 1}
	expected := 1

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}

func TestCase10(t *testing.T) {
	nums := []int{20}
	expected := 20

	actual := findMin(nums)

	assert.Equal(t, expected, actual)
}
