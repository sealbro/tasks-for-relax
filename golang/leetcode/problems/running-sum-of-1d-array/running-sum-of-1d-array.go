package running_sum_of_1d_array

// https://leetcode.com/problems/running-sum-of-1d-array/

func runningSum(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
}
