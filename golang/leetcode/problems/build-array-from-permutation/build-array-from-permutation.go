package build_array_from_permutation

// https://leetcode.com/problems/build-array-from-permutation/

func buildArray(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = nums[nums[i]]
	}

	return result
}
