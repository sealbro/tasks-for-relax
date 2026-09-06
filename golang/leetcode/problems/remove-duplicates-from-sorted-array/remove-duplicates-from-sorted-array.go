package remove_duplicates_from_sorted_array

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	ind := 1
	for i, num := range nums {
		if i != 0 && num != nums[i-1] {
			nums[ind] = num
			ind++
		}
	}

	return ind
}
