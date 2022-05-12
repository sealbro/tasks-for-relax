package find_first_and_last_position_of_element_in_sorted_array

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[0] > target || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}

	l1 := 0
	r2 := len(nums) - 1
	l2 := 0
	r1 := len(nums) - 1

	for l1 <= r1 || l2 <= r2 {
		p1 := l1 + (r1-l1)/2.0
		p2 := l2 + (r2-l2)/2.0

		if nums[p1] < target {
			l1 = p1 + 1
		} else {
			r1 = p1 - 1
		}

		if nums[p2] > target {
			r2 = p2 - 1
		} else {
			l2 = p2 + 1
		}
	}

	if l1 > -1 && r2 > -1 && nums[l1] == target && nums[r2] == target {
		return []int{l1, r2}
	}

	return []int{-1, -1}
}
