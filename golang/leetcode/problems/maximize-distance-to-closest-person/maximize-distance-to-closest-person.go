package maximize_distance_to_closest_person

// https://leetcode.com/problems/maximize-distance-to-closest-person/

func maxDistToClosest(seats []int) int {
	max := 0

	lastOccupied := -1

	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if lastOccupied == -1 {
				max = i
			} else {
				newMax := (i - lastOccupied) / 2

				if newMax > max {
					max = newMax
				}
			}
			lastOccupied = i
		}
	}

	if lastOccupied != len(seats)-1 {
		newMax := len(seats) - 1 - lastOccupied
		if newMax > max {
			max = newMax
		}
	}

	return max
}
