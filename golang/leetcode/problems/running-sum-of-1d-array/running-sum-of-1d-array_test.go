package running_sum_of_1d_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{1, 3, 6, 10}

	actual := runningSum(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{1, 1, 1, 1, 1}
	expected := []int{1, 2, 3, 4, 5}

	actual := runningSum(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{3, 1, 2, 10, 1}
	expected := []int{3, 4, 6, 16, 17}

	actual := runningSum(input)

	assert.EqualMany(t, expected, actual)
}
