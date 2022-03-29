package two_sum_ii_input_array_is_sorted

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{1, 2}
	actual := twoSum(nums, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{2, 3, 4}
	target := 6
	expected := []int{1, 3}
	actual := twoSum(nums, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{-1, 0}
	target := -1
	expected := []int{1, 2}
	actual := twoSum(nums, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{-1, 0, 7, 14}
	target := 13
	expected := []int{1, 4}
	actual := twoSum(nums, target)

	assert.EqualMany(t, expected, actual)
}
