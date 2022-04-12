package rotting_oranges

// https://leetcode.com/problems/rotting-oranges/

func orangesRotting(grid [][]int) int {
	iLength := len(grid)
	jLength := len(grid[0])

	findNum := 2
	multipleNum := -1
	freshFound := false
	changed := true

	times := 0
	for changed {
		changed = false
		freshFound = false
		for i := 0; i < iLength; i++ {
			for j := 0; j < jLength; j++ {
				if grid[i][j] == findNum {
					changed = rotting(grid, i, j, findNum*multipleNum, iLength, jLength) || changed
				}

				if grid[i][j] == 1 {
					freshFound = true
				}
			}
		}
		times++
		findNum = findNum * multipleNum
	}

	if freshFound {
		return -1
	}

	return times - 1
}

func rotting(grid [][]int, i, j int, value int, iLength, jLength int) bool {
	changed := false
	grid[i][j] = 99

	i0 := i-1 >= 0 && grid[i-1][j] == 1
	if i0 {
		grid[i-1][j] = value
		changed = true
	}

	j0 := j-1 >= 0 && grid[i][j-1] == 1
	if j0 {
		grid[i][j-1] = value
		changed = true
	}

	i1 := i+1 < iLength && grid[i+1][j] == 1
	if i1 {
		grid[i+1][j] = value
		changed = true
	}

	j1 := j+1 < jLength && grid[i][j+1] == 1
	if j1 {
		grid[i][j+1] = value
		changed = true
	}

	return changed
}
