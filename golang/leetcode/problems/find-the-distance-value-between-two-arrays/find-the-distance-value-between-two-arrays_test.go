package find_the_distance_value_between_two_arrays

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums1 := []int{4, 5, 8}
	nums2 := []int{10, 9, 1, 8}
	target := 2
	expected := 2

	actual := findTheDistanceValue(nums1, nums2, target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums1 := []int{1, 4, 2, 3}
	nums2 := []int{-4, -3, 6, 10, 20, 30}
	target := 3
	expected := 2

	actual := findTheDistanceValue(nums1, nums2, target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums1 := []int{2, 1, 100, 3}
	nums2 := []int{-5, -2, 10, -3, 7}
	target := 6
	expected := 1

	actual := findTheDistanceValue(nums1, nums2, target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums1 := []int{1, 10, 3, 20}
	nums2 := []int{10, 20, 30, 40}
	target := 6
	expected := 2

	actual := findTheDistanceValue(nums1, nums2, target)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums1 := []int{-803, 715, -224, 909, 121, -296, 872, 807, 715, 407, 94, -8, 572, 90, -520, -867, 485, -918, -827, -728, -653, -659, 865, 102, -564, -452, 554, -320, 229, 36, 722, -478, -247, -307, -304, -767, -404, -519, 776, 933, 236, 596, 954, 464}
	nums2 := []int{817, 1, -723, 187, 128, 577, -787, -344, -920, -168, -851, -222, 773, 614, -699, 696, -744, -302, -766, 259, 203, 601, 896, -226, -844, 168, 126, -542, 159, -833, 950, -454, -253, 824, -395, 155, 94, 894, -766, -63, 836, -433, -780, 611, -907, 695, -395, -975, 256, 373, -971, -813, -154, -765, 691, 812, 617, -919, -616, -510, 608, 201, -138, -669, -764, -77, -658, 394, -506, -675, 523, 730, -790, -109, 865, 975, -226, 651, 987, 111, 862, 675, -398, 126, -482, 457, -24, -356, -795, -575, 335, -350, -919, -945, -979, 611}
	target := 37
	expected := 0

	actual := findTheDistanceValue(nums1, nums2, target)

	assert.Equal(t, expected, actual)
}
