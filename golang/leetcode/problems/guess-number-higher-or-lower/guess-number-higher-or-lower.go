package guess_number_higher_or_lower

// https://leetcode.com/problems/guess-number-higher-or-lower/

var number = 0

func guessNumber(n int) int {
	start := 1
	pivot := n / 2
	end := n

	for eq := guess(pivot); eq != 0; eq = guess(pivot) {
		if eq == -1 {
			end, pivot = pivot, start+(pivot-start-1)/2
		} else if eq == 1 {
			start, pivot = pivot, pivot+(end-pivot+1)/2
		}
	}

	return pivot
}

func guess(num int) int {
	switch {
	case num > number:
		return -1
	case num < number:
		return 1
	default:
		return 0
	}
}
