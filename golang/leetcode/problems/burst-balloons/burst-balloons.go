package burst_balloons

import (
	"math"
)

// https://leetcode.com/problems/burst-balloons/

func maxCoins(nums []int) int {
	n := len(nums) + 2
	newNums := make([]int, n)
	for i := 1; i < len(nums)+1; i++ {
		newNums[i] = nums[i-1]
	}

	newNums[0] = 1
	newNums[n-1] = 1

	var dp [][]int
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, n))
	}

	for left := n - 2; left >= 1; left-- {
		for right := left; right <= n-2; right++ {
			for i := left; i <= right; i++ {
				gain := newNums[left-1] * newNums[i] * newNums[right+1]
				remaining := dp[left][i-1] + dp[i+1][right]
				dp[left][right] = int(math.Max(float64(remaining+gain), float64(dp[left][right])))
			}
		}
	}

	return dp[1][n-2]
}
