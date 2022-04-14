package maximum_subarray

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6

	actual := maxSubArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1}
	expected := 1

	actual := maxSubArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}
	expected := 23

	actual := maxSubArray(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{1, 2}
	expected := 3

	actual := maxSubArray(nums)

	assert.Equal(t, expected, actual)
}
