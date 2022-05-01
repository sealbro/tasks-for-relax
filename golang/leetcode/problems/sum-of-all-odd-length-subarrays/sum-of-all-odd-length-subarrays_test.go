package sum_of_all_odd_length_subarrays

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 4, 2, 5, 3}
	expected := 58

	actual := sumOddLengthSubarrays(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2}
	expected := 3

	actual := sumOddLengthSubarrays(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{10, 11, 12}
	expected := 66

	actual := sumOddLengthSubarrays(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{1}
	expected := 1

	actual := sumOddLengthSubarrays(nums)

	assert.Equal(t, expected, actual)
}
