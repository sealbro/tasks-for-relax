package pascals_triangle

// https://leetcode.com/problems/pascals-triangle/

func generate(numRows int) [][]int {
	pascalsTriangle := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		lastIndex := len(row) - 1
		row[lastIndex] = 1

		rowIndex := i - 1
		for j := 1; j < lastIndex; j++ {
			left := 0
			right := pascalsTriangle[rowIndex][j]

			if j-1 >= 0 {
				left = pascalsTriangle[rowIndex][j-1]
			}

			row[j] = left + right
		}

		pascalsTriangle = append(pascalsTriangle, row)
	}

	return pascalsTriangle
}
