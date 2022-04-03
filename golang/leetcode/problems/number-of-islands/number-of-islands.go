package number_of_islands

// https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	rowsLength := len(grid)
	colLength := len(grid[0])

	count := 0

	for r := 0; r < rowsLength; r++ {
		for c := 0; c < colLength; c++ {
			if grid[r][c] == 49 {
				count++
			} else {
				continue
			}

			findAndSet(grid, r, c, rowsLength, colLength)
		}
	}

	return count
}

func findAndSet(grid [][]byte, sr int, sc int, rowsLength, colLength int) {

	switch grid[sr][sc] {
	case 49:
		grid[sr][sc] -= 49
	default:
		return
	}

	if sr-1 >= 0 {
		findAndSet(grid, sr-1, sc, rowsLength, colLength)
	}

	if sr+1 < rowsLength {
		findAndSet(grid, sr+1, sc, rowsLength, colLength)
	}

	if sc-1 >= 0 {
		findAndSet(grid, sr, sc-1, rowsLength, colLength)
	}

	if sc+1 < colLength {
		findAndSet(grid, sr, sc+1, rowsLength, colLength)
	}
}
