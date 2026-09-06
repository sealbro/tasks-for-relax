package remove_duplicates_from_sorted_array_ii

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

func removeDuplicates(nums []int) int {
	ind := 1
	last := nums[0]
	one := true
	for i := 1; i < len(nums); i++ {
		if last == nums[i] && one {
			nums[ind] = nums[i]
			ind++
			one = false
			continue
		}
		if last != nums[i] {
			nums[ind] = nums[i]
			last = nums[i]
			one = true
			ind++
		}
	}

	return ind
}
