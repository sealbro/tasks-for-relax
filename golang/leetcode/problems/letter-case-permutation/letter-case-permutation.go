package letter_case_permutation

import (
	"math"
)

// https://leetcode.com/problems/letter-case-permutation/

func letterCasePermutation(s string) []string {
	var symbols = []byte(s)
	var output []string

	var indexes []int
	for i := 0; i < len(symbols); i++ {
		if symbols[i] > 64 {
			indexes = append(indexes, i)
		}
	}

	lettersLength := len(indexes)
	count := int(math.Pow(2, float64(lettersLength)))
	for i := 0; i < count; i++ {
		for j := 0; j < lettersLength; j++ {
			q := (i >> j) & 1

			if q == 1 && symbols[indexes[j]] > 96 {
				symbols[indexes[j]] = symbols[indexes[j]] - 32
				continue
			}

			if q == 0 && symbols[indexes[j]] < 96 {
				symbols[indexes[j]] = symbols[indexes[j]] + 32
				continue
			}
		}
		output = append(output, string(symbols))
	}

	return output
}
