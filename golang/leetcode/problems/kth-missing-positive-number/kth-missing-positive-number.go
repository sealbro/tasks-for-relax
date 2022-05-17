package kth_missing_positive_number

// https://leetcode.com/problems/kth-missing-positive-number/

func findKthPositive(arr []int, k int) int {
	l := 0
	r := len(arr) - 1

	left := 1
	numbersBefore := 1

	if k <= arr[0]-1 {
	} else if k > arr[r]-(r+1) {
		numbersBefore = arr[r] - (r + 1)
		left = arr[r]
	} else {
		for l <= r {
			p := l + (r-l)/2.0

			if arr[p]-(p+1) < k {
				l = p + 1
			} else {
				r = p - 1
			}
		}

		left = arr[r]
		numbersBefore = arr[r] - (r + 1)
	}

	return left + (k - numbersBefore)
}
