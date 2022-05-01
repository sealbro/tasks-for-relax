package richest_customer_wealth

// https://leetcode.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
	maxSum := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if maxSum < sum {
			maxSum = sum
		}
	}

	return maxSum
}
