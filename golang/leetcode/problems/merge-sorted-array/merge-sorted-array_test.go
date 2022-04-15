package merge_sorted_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	expected := []int{1, 2, 2, 3, 5, 6}

	merge(nums1, 3, nums2, 3)
	actual := nums1

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums1 := []int{0}
	nums2 := []int{1}
	expected := []int{1}

	merge(nums1, 0, nums2, 1)
	actual := nums1

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums1 := []int{1}
	var nums2 []int
	expected := []int{1}

	merge(nums1, 1, nums2, 0)
	actual := nums1

	assert.EqualMany(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums1 := []int{0}
	var nums2 []int
	expected := []int{0}

	merge(nums1, 0, nums2, 0)
	actual := nums1

	assert.EqualMany(t, expected, actual)
}

func TestCase6(t *testing.T) {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	expected := []int{1, 2}

	merge(nums1, 1, nums2, 1)
	actual := nums1

	assert.EqualMany(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	expected := []int{1, 2, 2, 3, 4, 5, 6}

	merge(nums1, 4, nums2, 3)
	actual := nums1

	assert.EqualMany(t, expected, actual)
}
