package largest_perimeter_triangle

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 1, 2}
	expected := 5

	actual := largestPerimeter(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 1}
	expected := 0

	actual := largestPerimeter(nums)

	assert.Equal(t, expected, actual)
}
