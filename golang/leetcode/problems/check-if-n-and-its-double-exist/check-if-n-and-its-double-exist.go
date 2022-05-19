package check_if_n_and_its_double_exist

import "sort"

// https://leetcode.com/problems/check-if-n-and-its-double-exist/

func checkIfExist(arr []int) bool {
	sort.Ints(arr)

	for i, val := range arr {
		if findNum(arr, i, val*2) {
			return true
		}
	}

	return false
}

func findNum(arr []int, i, num int) bool {
	l := 0
	r := len(arr) - 1

	for l <= r {
		p := l + (r-l)/2

		if p != i && arr[p] == num {
			return true
		}

		if arr[p] < num {
			l = p + 1
		} else {
			r = p - 1
		}
	}

	return false
}
