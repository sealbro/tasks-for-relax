package search_in_a_sorted_array_of_unknown_size

import "math"

// https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/

type ArrayReader struct {
	bucket []int
}

func (a *ArrayReader) get(index int) int {
	if index < 0 || index > len(a.bucket)-1 {
		return math.MaxInt32
	}

	return a.bucket[index]
}

func search(reader ArrayReader, target int) int {
	var pivot int
	findLeft := 0
	findRight := int(math.Pow(10, 4)) // constrains max index

	for findLeft+1 < findRight {
		pivot = findLeft + (findRight-findLeft)/2

		if reader.get(pivot) == math.MaxInt32 {
			findRight = pivot
		} else {
			findLeft = pivot
		}
	}

	pivot = 0
	left := 0
	right := findLeft

	for left <= right {
		pivot = left + (right-left)/2

		if reader.get(pivot) == target {
			return pivot
		}

		if reader.get(pivot) < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return -1
}
