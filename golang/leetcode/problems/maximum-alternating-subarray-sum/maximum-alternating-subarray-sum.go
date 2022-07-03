package maximum_alternating_subarray_sum

import "math"

// https://leetcode.com/problems/maximum-alternating-subarray-sum/

func maximumAlternatingSubarraySum(nums []int) int64 {
	length := len(nums)

	if length == 1 {
		return int64(nums[0])
	}

	var max = float64(nums[0])

	first := float64(nums[0])
	second := float64(0)
	for i := 1; i < length; i++ {
		if i%2 == 0 {
			first = math.Max(float64(nums[i])+first, float64(nums[i]))
			second = second - float64(nums[i])
		} else {
			first = first - float64(nums[i])
			second = math.Max(second+float64(nums[i]), float64(nums[i]))
		}
		max = math.Max(max, math.Max(first, second))
	}
	return int64(max)
}
