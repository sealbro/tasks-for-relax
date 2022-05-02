package matrix_diagonal_sum

// https://leetcode.com/problems/matrix-diagonal-sum/

func diagonalSum(mat [][]int) int {
	sum := 0
	length := len(mat)
	for i := 0; i < length; i++ {
		sum += mat[i][i] + mat[i][length-i-1]
	}

	if length%2 != 0 {
		mid := length / 2
		sum -= mat[mid][mid]
	}

	return sum
}
