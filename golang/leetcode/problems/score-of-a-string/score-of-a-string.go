package score_of_a_string

import "math"

// https://leetcode.com/problems/score-of-a-string/

func scoreOfString(s string) int {

	sum := 0
	for i := 0; i < len(s)-1; i++ {
		sum += int(math.Abs(float64(s[i]) - float64(s[i+1])))
	}

	return sum
}
