package search_in_rotated_sorted_array

// https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	if len(nums) == 1 {
		switch {
		case nums[0] == target:
			return 0
		default:
			return -1
		}
	}

	first := nums[0]
	last := nums[len(nums)-1]

	if first > target && last < target {
		return -1
	}

	switch target {
	case first:
		return 0
	case last:
		return len(nums) - 1
	}

	l := 0
	r := len(nums) - 1
	for l <= r {
		p := l + (r-l)/2

		if nums[p] == target {
			return p
		}

		isReverse := (target < first && nums[p] > last && nums[p] > target) ||
			(target > last && nums[p] < first && nums[p] < target)

		if (nums[p] < target && !isReverse) || (nums[p] > target && isReverse) {
			l = p + 1
		} else {
			r = p - 1
		}
	}

	return -1
}
