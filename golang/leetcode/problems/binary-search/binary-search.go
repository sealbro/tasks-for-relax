package binary_search

// https://leetcode.com/problems/binary-search/solution/

func search(nums []int, target int) int {
	var pivot int
	left := 0
	right := len(nums) - 1

	for {
		if left > right {
			break
		}

		pivot = left + (right-left)/2

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return -1
}
