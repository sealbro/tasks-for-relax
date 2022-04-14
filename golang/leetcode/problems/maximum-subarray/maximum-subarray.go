package maximum_subarray

// https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum+nums[i] < nums[i] {
			currentSum = nums[i]
		} else {
			currentSum = currentSum + nums[i]
		}

		if maxSum < currentSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
