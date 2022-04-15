package merge_sorted_array

// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for i := 0; i < len(nums1); i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	lastX := len(nums1) - n - 1
	lastY := n - 1
	for i := len(nums1) - 1; i >= 0; i-- {
		if lastY < 0 || lastX >= 0 && nums1[lastX] > nums2[lastY] {
			nums1[i] = nums1[lastX]
			lastX--
			continue
		} else {
			nums1[i] = nums2[lastY]
			lastY--
		}
	}
}
