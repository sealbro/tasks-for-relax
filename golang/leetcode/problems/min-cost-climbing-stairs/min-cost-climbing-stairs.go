package min_cost_climbing_stairs

import "math"

// https://leetcode.com/problems/min-cost-climbing-stairs/

func minCostClimbingStairs(cost []int) int {
	one, two := 0, 0
	for i := 2; i < len(cost)+1; i++ {
		temp := one
		one = int(math.Min(float64(one+cost[i-1]), float64(two+cost[i-2])))
		two = temp
	}

	return one
}
