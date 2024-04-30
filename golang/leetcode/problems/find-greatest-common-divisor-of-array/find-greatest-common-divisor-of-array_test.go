package find_greatest_common_divisor_of_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 5, 6, 9, 10}
	expected := 2

	actual := findGCD(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{7, 5, 6, 8, 3}
	expected := 1

	actual := findGCD(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{3, 3}
	expected := 3

	actual := findGCD(nums)

	assert.Equal(t, expected, actual)
}
