package search_a_2d_matrix

// https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	if matrix[0][0] == target {
		return true
	}

	min := 0
	max := len(matrix) - 1
	pivot := max / 2

	for min < max {
		if matrix[pivot][0] == target {
			return true
		}

		if matrix[pivot][0] > target {
			max = pivot - 1
		} else {
			min = pivot
		}

		pivot = min + ((max - min) / 2)

		if min == pivot {
			pivot = max
		}
	}

	row := min
	min = 0
	max = len(matrix[row]) - 1
	pivot = max / 2

	for min < max {
		if matrix[row][pivot] == target {
			return true
		}

		if matrix[row][pivot] > target {
			max = pivot - 1
		} else {
			min = pivot
		}

		pivot = min + ((max - min) / 2)

		if min == pivot {
			pivot = max
		}
	}

	return false
}
