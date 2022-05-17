package arranging_coins

// https://leetcode.com/problems/arranging-coins/

func arrangeCoins(n int) int {
	l := 0
	r := n
	c := 0
	for l <= r {
		p := l + (r-l)/2
		c = p * (p + 1) / 2

		if c == n {
			return p
		}

		if n < c {
			r = p - 1
		} else {
			l = p + 1
		}
	}

	return r
}
