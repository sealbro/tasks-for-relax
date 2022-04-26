package count_odd_numbers_in_an_interval_range

// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/

func countOdds(low int, high int) int {
	i := (high - low) / 2

	if low&1 == 1 || high&1 == 1 {
		i++
	}

	return i
}
