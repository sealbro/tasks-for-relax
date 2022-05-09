package peak_index_in_a_mountain_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{0, 1, 0}
	expected := 1
	actual := peakIndexInMountainArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{0, 2, 1, 0}
	expected := 1
	actual := peakIndexInMountainArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{0, 10, 5, 2}
	expected := 1
	actual := peakIndexInMountainArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{3, 4, 5, 1}
	expected := 2
	actual := peakIndexInMountainArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums := []int{3, 7, 5, 1}
	expected := 1
	actual := peakIndexInMountainArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	nums := []int{40, 48, 61, 75, 100, 99, 98, 39, 30, 10}
	expected := 4
	actual := peakIndexInMountainArray(nums)

	assert.Equal(t, expected, actual)
}
