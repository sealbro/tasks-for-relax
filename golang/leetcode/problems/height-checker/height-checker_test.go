package height_checker

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 1, 4, 2, 1, 3}
	expected := 3

	actual := heightChecker(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{5, 1, 2, 3, 4}
	expected := 5

	actual := heightChecker(nums)

	assert.Equal(t, expected, actual)
}
