package reshape_the_matrix

// https://leetcode.com/problems/reshape-the-matrix/

func matrixReshape(mat [][]int, r int, c int) [][]int {
	var reshaped [][]int

	cols := len(mat[0])
	rows := len(mat)
	if rows*cols != r*c && cols != c {
		c = cols
	}

	nr := 0
	nc := 0
	for i := 0; i < rows; i++ {

		for j := 0; j < cols; j++ {
			if nr < r && nc == c {
				nc = 0
				nr++
			}

			if len(reshaped) < nr+1 {
				reshaped = append(reshaped, []int{})
			}

			if nc < c && nr < r {
				reshaped[nr] = append(reshaped[nr], mat[i][j])
				nc++
			}

		}
	}

	return reshaped
}
