package intersection_of_two_arrays_ii

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2, 2}

	actual := intersect(nums1, nums2)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	expected := []int{4, 9}

	actual := intersect(nums1, nums2)

	assert.EqualMany(t, expected, actual)
}
