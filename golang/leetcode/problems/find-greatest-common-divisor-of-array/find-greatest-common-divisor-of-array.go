package find_greatest_common_divisor_of_array

// https://leetcode.com/problems/find-greatest-common-divisor-of-array/

func findGCD(nums []int) int {
	minV := nums[0]
	maxV := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}

		if nums[i] < minV {
			minV = nums[i]
		}
	}

	for maxV != 0 {
		t := maxV
		maxV = minV % maxV
		minV = t
	}

	return minV
}
