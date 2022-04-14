package power_of_two

// https://leetcode.com/problems/power-of-two/

func isPowerOfTwo(n int) bool {
	countBits := 0
	for i := 0; i < 33; i++ {
		bit := n >> i & 1

		if bit == 1 {
			countBits++

			if countBits > 1 {
				return false
			}
		}
	}

	return countBits == 1
}
