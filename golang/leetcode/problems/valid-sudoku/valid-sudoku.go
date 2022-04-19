package valid_sudoku

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	colRow := 0

	const size = 9

	for colRow < size {
		digits := [size]int{}
		for i := 0; i < size; i++ {
			if board[i][colRow] != '.' {
				digit := board[i][colRow] - '1'
				if digits[digit] == 1 {
					return false
				}
				digits[digit] += 1
			}
		}

		digits = [size]int{}
		for i := 0; i < size; i++ {
			if board[colRow][i] != '.' {
				digit := board[colRow][i] - '1'
				if digits[digit] == 1 {
					return false
				}
				digits[digit] += 1
			}
		}
		colRow++
	}

	for row := 0; row < size; row += 3 {
		for col := 0; col < size; col += 3 {
			digits := [size]int{}
			for i := row; i < row+3; i++ {
				for j := col; j < col+3; j++ {
					if board[i][j] != '.' {
						digit := board[i][j] - '1'
						if digits[digit] == 1 {
							return false
						}
						digits[digit] += 1
					}
				}
			}
		}
	}

	return true
}
