package find_the_distance_value_between_two_arrays

import (
	"sort"
)

// https://leetcode.com/problems/find-the-distance-value-between-two-arrays/

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	count := 0

	for _, value := range arr1 {
		if isValid(arr2, value-d, value+d) {
			count++
		}
	}

	return count
}

func isValid(arr2 []int, left, right int) bool {
	l := 0
	r := len(arr2) - 1

	for l <= r {
		p := l + (r-l)/2
		if arr2[p] >= left && arr2[p] <= right {
			return false
		}

		if arr2[p] < left {
			l = p + 1
		} else {
			r = p - 1
		}
	}

	return true
}
