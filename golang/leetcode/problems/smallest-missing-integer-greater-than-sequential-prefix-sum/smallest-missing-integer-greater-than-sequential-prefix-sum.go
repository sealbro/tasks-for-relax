package smallest_missing_integer_greater_than_sequential_prefix_sum

import "slices"

// https://leetcode.com/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/

func missingInteger(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0] + 1
	}

	last := nums[0]
	sum := last

	for i, n := range nums[1:] {
		if last+1 == n {
			last = n
			sum += n
		} else {
			sorted := nums[i:]
			slices.Sort(sorted)
			for j := 0; j < len(sorted); j++ {
				if sorted[j] == sum {
					sum++
				}
			}
			break
		}
	}

	return sum
}
