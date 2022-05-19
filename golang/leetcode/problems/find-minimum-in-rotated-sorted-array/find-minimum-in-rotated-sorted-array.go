package find_minimum_in_rotated_sorted_array

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

func findMin(nums []int) int {
	if nums[len(nums)-1] >= nums[0] {
		return nums[0]
	}

	first := nums[0]
	//last := nums[len(nums)-1]

	l := 0
	r := len(nums) - 1
	for l <= r {
		p := l + (r-l)/2

		isReverse := nums[p] < first

		if nums[p] >= first && !isReverse {
			l = p + 1
		} else {
			r = p - 1
		}
	}

	return nums[l]
}
