package merge_strings_alternately

import (
	"math"
	"strings"
)

// https://leetcode.com/problems/merge-strings-alternately/

func mergeAlternately(word1 string, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)
	maxLength := int(math.Max(float64(len1), float64(len2)))

	builder := strings.Builder{}

	for i := 0; i < maxLength; i++ {
		if i < len1 {
			builder.WriteByte(word1[i])
		}
		if i < len2 {
			builder.WriteByte(word2[i])
		}
	}

	return builder.String()
}
