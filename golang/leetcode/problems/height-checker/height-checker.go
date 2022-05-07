package height_checker

import "sort"

// https://leetcode.com/problems/height-checker/

func heightChecker(heights []int) int {
	var copyHeights = make([]int, len(heights))
	copy(copyHeights, heights)
	sort.Ints(copyHeights)

	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != copyHeights[i] {
			count++
		}
	}

	return count
}
