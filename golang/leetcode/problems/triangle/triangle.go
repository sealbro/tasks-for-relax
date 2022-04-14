package triangle

import "math"

// https://leetcode.com/problems/triangle/

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	minResult := math.MaxInt32
	maxLevel := len(triangle)
	for i := 1; i < maxLevel; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = triangle[i][j] + min(triangle, i-1, j)

			if i == maxLevel-1 {
				if triangle[i][j] < minResult {
					minResult = triangle[i][j]
				}
			}
		}
	}

	return minResult
}

func min(triangle [][]int, level, index int) int {
	if level < 0 {
		return 0
	}

	fi := index - 1
	li := index

	if fi < 0 {
		fi = index
	}

	if li >= len(triangle[level])-1 {
		li = len(triangle[level]) - 1
	}

	minResult := math.MaxInt32
	for i := fi; i < li+1; i++ {
		if triangle[level][i] < minResult {
			minResult = triangle[level][i]
		}
	}

	return minResult
}
