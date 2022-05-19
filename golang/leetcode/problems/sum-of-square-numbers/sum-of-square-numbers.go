package sum_of_square_numbers

// https://leetcode.com/problems/sum-of-square-numbers/

func judgeSquareSum(c int) bool {
	var a int64 = 0
	c64 := int64(c)
	for ; a*a <= c64; a++ {
		b := c64 - (a * a)
		if binarySearch(0, b, b) {
			return true
		}
	}
	return false
}

func binarySearch(s, e, n int64) bool {
	if s > e {
		return false
	}

	var mid = s + (e-s)/2
	if mid*mid == n {
		return true
	}
	if mid*mid > n {
		return binarySearch(s, mid-1, n)
	}

	return binarySearch(mid+1, e, n)
}
