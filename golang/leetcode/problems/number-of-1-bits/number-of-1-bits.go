package number_of_1_bits

// https://leetcode.com/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
	var count uint32 = 0
	for i := 0; i < 33; i++ {
		count += num >> i & 1
	}

	return int(count)
}
