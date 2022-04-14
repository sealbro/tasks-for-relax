package house_robber

import (
	"math"
)

// https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
	lastSum := 0
	expectedSum := nums[0]
	for i := 1; i < len(nums); i++ {
		lastSum, expectedSum = expectedSum, int(math.Max(float64(expectedSum), float64(lastSum+nums[i])))
	}

	return expectedSum
}
