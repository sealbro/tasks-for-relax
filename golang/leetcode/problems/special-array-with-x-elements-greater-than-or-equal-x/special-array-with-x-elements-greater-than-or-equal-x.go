package special_array_with_x_elements_greater_than_or_equal_x

import "sort"

// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/

func specialArray(nums []int) int {
	sort.Ints(nums)

	l := 0
	r := len(nums) - 1

	if len(nums) == 1 && nums[0] == 0 {
		return -1
	}

	if len(nums) <= nums[0] {
		return len(nums)
	}

	for l <= r {
		p := l + (r-l)/2

		if len(nums)-(p+1) <= nums[p] {
			r = p - 1
		} else {
			l = p + 1
		}
	}

	diff := len(nums) - l
	if nums[r] < diff && diff <= nums[l] {
		return diff
	}

	if nums[l] > 0 && nums[l] >= diff {
		return nums[l]
	}

	return -1
}
