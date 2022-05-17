package special_array_with_x_elements_greater_than_or_equal_x

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{3, 5}
	expected := 2

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{0, 4, 3, 0, 4}
	expected := 3

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{0, 0}
	expected := -1

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{0, 0, 0, 0, 0, 0, 0}
	expected := -1

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums := []int{0, 0, 0, 0, 5, 6, 6}
	expected := 3

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	nums := []int{0, 0, 0, 4, 5, 6, 6}
	expected := 4

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase7(t *testing.T) {
	nums := []int{1, 0, 0, 6, 4, 9}
	expected := 3

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase8(t *testing.T) {
	nums := []int{1, 0, 0, 6, 2, 9}
	expected := -1

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase9(t *testing.T) {
	nums := []int{0, 5, 0, 1, 8, 3, 0, 1}
	expected := 3

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase10(t *testing.T) {
	nums := []int{0}
	expected := -1

	actual := specialArray(nums)

	assert.Equal(t, expected, actual)
}
