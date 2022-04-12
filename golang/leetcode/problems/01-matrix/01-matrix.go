package n01_matrix

// https://leetcode.com/problems/01-matrix/

func updateMatrix(mat [][]int) [][]int {
	mLength := len(mat)
	nLength := len(mat[0])

	calcCount := 1
	for i := 0; i < calcCount; i++ {
		for mi := 0; mi < mLength; mi++ {
			for ni := 0; ni < nLength; ni++ {
				if mat[mi][ni] == 0 {
					continue
				}
				minValue := min(mat, mi, ni, mLength, nLength)
				mat[mi][ni] = minValue

				if calcCount < minValue {
					calcCount = minValue
				}
			}
		}
	}

	return mat
}

func min(mat [][]int, mi int, ni int, mLength int, nLength int) int {
	minValue := -1

	if mi-1 >= 0 {
		value := mat[mi-1][ni]
		if value < minValue {
			minValue = value
		}
	}

	if mi+1 < mLength {
		value := mat[mi+1][ni]
		if minValue == -1 || value < minValue {
			minValue = value
		}
	}

	if ni-1 >= 0 {
		value := mat[mi][ni-1]
		if minValue == -1 || value < minValue {
			minValue = value
		}
	}

	if ni+1 < nLength {
		value := mat[mi][ni+1]
		if minValue == -1 || value < minValue {
			minValue = value
		}
	}

	if mat[mi][ni] <= minValue {
		return minValue + 1
	}

	return mat[mi][ni]
}
