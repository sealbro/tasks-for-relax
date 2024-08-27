package check_if_bitwise_or_has_trailing_zeros

// https://leetcode.com/problems/check-if-bitwise-or-has-trailing-zeros/

func hasTrailingZeros(nums []int) bool {
	count := 0
	for _, n := range nums {
		if n&1 == 1 {
			continue
		}

		count++
		if count == 2 {
			return true
		}
	}

	return false
}
