package count_prefix_and_suffix_pairs_i

import "strings"

// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/

func countPrefixSuffixPairs(words []string) int {
	if len(words) == 1 {
		return 0
	}

	count := 0

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.HasPrefix(words[j], words[i]) && strings.HasSuffix(words[j], words[i]) {
				count++
			}
		}
	}

	return count
}
