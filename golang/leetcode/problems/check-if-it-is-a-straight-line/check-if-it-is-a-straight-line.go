package check_if_it_is_a_straight_line

// https://leetcode.com/problems/check-if-it-is-a-straight-line/

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}

	for i := 2; i < len(coordinates); i++ {
		if !checkPoint(coordinates, coordinates[i][0], coordinates[i][1]) {
			return false
		}
	}

	return true
}

func checkPoint(coordinates [][]int, x, y int) bool {
	y2 := coordinates[1][1]
	y1 := coordinates[0][1]
	x2 := coordinates[1][0]
	x1 := coordinates[0][0]

	return ((x1 - x2) * (y1 - y)) == ((x1 - x) * (y1 - y2))
}
