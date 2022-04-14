package single_number

// https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
	xor := 0
	for _, num := range nums {
		xor = xor ^ num
	}

	return xor
}
