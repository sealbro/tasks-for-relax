package search_insert_position

// https://leetcode.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	var pivot int
	left := 0
	right := len(nums) - 1

	for left <= right {
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

	return left
}
