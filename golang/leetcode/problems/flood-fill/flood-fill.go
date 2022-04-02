package flood_fill

// https://leetcode.com/problems/flood-fill/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]

	rowsLength := len(image)
	colLength := len(image[0])

	if oldColor != newColor {
		findAndSet(image, sr, sc, newColor, oldColor, rowsLength, colLength)
	}

	return image
}

func findAndSet(image [][]int, sr int, sc int, newColor, oldColor int, rowsLength, colLength int) {
	if image[sr][sc] == oldColor {
		image[sr][sc] = newColor
	}

	if sr-1 >= 0 && image[sr-1][sc] == oldColor {
		findAndSet(image, sr-1, sc, newColor, oldColor, rowsLength, colLength)
	}

	if sr+1 < rowsLength && image[sr+1][sc] == oldColor {
		findAndSet(image, sr+1, sc, newColor, oldColor, rowsLength, colLength)
	}

	if sc-1 >= 0 && image[sr][sc-1] == oldColor {
		findAndSet(image, sr, sc-1, newColor, oldColor, rowsLength, colLength)
	}

	if sc+1 < colLength && image[sr][sc+1] == oldColor {
		findAndSet(image, sr, sc+1, newColor, oldColor, rowsLength, colLength)
	}
}
