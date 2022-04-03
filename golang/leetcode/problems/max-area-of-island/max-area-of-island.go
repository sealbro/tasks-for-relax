package max_area_of_island

// https://leetcode.com/problems/max-area-of-island/

func maxAreaOfIsland(grid [][]int) int {
	rowsLength := len(grid)
	colLength := len(grid[0])

	max := 0

	for r := 0; r < rowsLength; r++ {
		for c := 0; c < colLength; c++ {
			if grid[r][c] == 0 {
				continue
			}
			count := 0
			findAndSet(grid, r, c, rowsLength, colLength, &count)
			if max < count {
				max = count
			}
		}
	}

	return max
}

func findAndSet(grid [][]int, sr int, sc int, rowsLength, colLength int, count *int) {

	switch grid[sr][sc] {
	case 1:
		*count++
		grid[sr][sc] -= 2
	default:
		return
	}

	if sr-1 >= 0 {
		findAndSet(grid, sr-1, sc, rowsLength, colLength, count)
	}

	if sr+1 < rowsLength {
		findAndSet(grid, sr+1, sc, rowsLength, colLength, count)
	}

	if sc-1 >= 0 {
		findAndSet(grid, sr, sc-1, rowsLength, colLength, count)
	}

	if sc+1 < colLength {
		findAndSet(grid, sr, sc+1, rowsLength, colLength, count)
	}
}
