package count_negative_numbers_in_a_sorted_matrix

// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/

func countNegatives(grid [][]int) int {
	count := 0

	for _, row := range grid {
		count += countNeg(row)
	}

	return count
}

func countNeg(arr []int) int {
	if arr[0] < 0 {
		return len(arr)
	}

	if len(arr) == 1 && arr[0] >= 0 {
		return 0
	}

	l := 0
	r := len(arr) - 1

	for l <= r {
		p := l + (r-l)/2

		if arr[p] < 0 {
			r = p - 1
		} else {
			l = p + 1
		}
	}

	if arr[r] > -1 && l >= len(arr) || arr[l] > -1 {
		return 0
	}

	return len(arr) - l
}
