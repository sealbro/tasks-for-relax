package maximize_distance_to_closest_person

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 0, 0, 0, 1, 0, 1}
	expected := 2

	actual := maxDistToClosest(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{0, 1}
	expected := 1

	actual := maxDistToClosest(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{1, 0, 0, 0}
	expected := 3

	actual := maxDistToClosest(nums)

	assert.Equal(t, expected, actual)
}
