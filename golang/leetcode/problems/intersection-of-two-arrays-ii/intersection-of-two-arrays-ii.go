package intersection_of_two_arrays_ii

// https://leetcode.com/problems/intersection-of-two-arrays-ii/

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	var intersection []int
	var nums = make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		nums[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		if val, ok := nums[nums2[i]]; ok && val != 0 {
			nums[nums2[i]]--
			intersection = append(intersection, nums2[i])
		}
	}

	return intersection
}
