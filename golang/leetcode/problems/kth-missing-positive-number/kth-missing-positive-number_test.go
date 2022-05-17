package kth_missing_positive_number

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 3, 4, 7, 11}
	target := 5
	expected := 9

	actual := findKthPositive(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	target := 2
	expected := 6

	actual := findKthPositive(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{3, 4}
	target := 1
	expected := 1

	actual := findKthPositive(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{3}
	target := 6
	expected := 7

	actual := findKthPositive(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums := []int{2, 3, 4, 7, 11}
	target := 7
	expected := 12

	actual := findKthPositive(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	nums := []int{1, 3}
	target := 1
	expected := 2

	actual := findKthPositive(nums, target)

	assert.Equal(t, expected, actual)
}
