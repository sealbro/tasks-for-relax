package maximum_distance_between_a_pair_of_values

// https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/

func maxDistance(nums1 []int, nums2 []int) int {
	maxDist := 0

	for i, num1 := range nums1 {
		l := 0
		r := len(nums2) - 1
		for l <= r {
			p := l + (r-l)/2

			if num1 <= nums2[p] {
				l = p + 1
			} else {
				r = p - 1
			}
		}

		if i <= l && r-i > maxDist {
			maxDist = r - i
		}
	}

	return maxDist
}
