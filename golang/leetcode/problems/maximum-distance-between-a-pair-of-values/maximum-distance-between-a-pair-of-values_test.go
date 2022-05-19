package maximum_distance_between_a_pair_of_values

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums1 := []int{55, 30, 5, 4, 2}
	nums2 := []int{100, 20, 10, 10, 5}

	expected := 2

	actual := maxDistance(nums1, nums2)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums1 := []int{2, 2, 2}
	nums2 := []int{10, 10, 1}

	expected := 1

	actual := maxDistance(nums1, nums2)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums1 := []int{30, 29, 19, 5}
	nums2 := []int{25, 25, 25, 25, 25}

	expected := 2

	actual := maxDistance(nums1, nums2)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums1 := []int{5, 4, 3, 2}
	nums2 := []int{1, 1, 1, 1, 1}

	expected := 0

	actual := maxDistance(nums1, nums2)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums1 := []int{2}
	nums2 := []int{5}

	expected := 0

	actual := maxDistance(nums1, nums2)

	assert.Equal(t, expected, actual)
}
