package happy_number

import (
	"math"
)

// https://leetcode.com/problems/happy-number/

func isHappy(n int) bool {
	hashMap := make(map[int]struct{})
	for n != 1 {
		currentSum := 0
		for n > 0 {
			currentSum += int(math.Pow(float64(n%10), 2.0))
			n = n / 10
		}
		n = currentSum

		if _, ok := hashMap[n]; ok {
			return false
		}

		hashMap[n] = struct{}{}
	}

	return true
}
