package count_prefix_and_suffix_pairs_i

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{input: []string{"a", "aba", "ababa", "aa"}, expected: 4},
		{input: []string{"pa", "papa", "ma", "mama"}, expected: 2},
		{input: []string{"abab", "ab"}, expected: 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := countPrefixSuffixPairs(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
