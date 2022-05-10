package valid_perfect_square

// https://leetcode.com/problems/valid-perfect-square/

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	pivot := num / 2
	for pivot*pivot > num {
		pivot = (pivot + (num / pivot)) / 2
	}

	return pivot*pivot == num
}
