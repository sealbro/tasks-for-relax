package find_nearest_point_that_has_the_same_x_or_y_coordinate

import "math"

// https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/

func nearestValidPoint(x int, y int, points [][]int) int {
	var diff = math.MaxFloat32
	index := -1
	for i := 0; i < len(points); i++ {
		if x == points[i][0] || y == points[i][1] {
			newDiff := math.Abs(float64(x-points[i][0])) + math.Abs(float64(y-points[i][1]))
			if newDiff < diff {
				diff = newDiff
				index = i
			}
		}
	}

	return index
}
