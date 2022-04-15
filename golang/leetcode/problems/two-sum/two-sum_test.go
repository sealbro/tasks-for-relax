package two_sum

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}

	actual := twoSum(nums, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}

	actual := twoSum(nums, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expected := []int{0, 1}

	actual := twoSum(nums, target)

	assert.EqualMany(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{2, 7, 3, 1, 2}
	target := 4
	expected := []int{2, 3}

	actual := twoSum(nums, target)

	assert.EqualMany(t, expected, actual)
}
