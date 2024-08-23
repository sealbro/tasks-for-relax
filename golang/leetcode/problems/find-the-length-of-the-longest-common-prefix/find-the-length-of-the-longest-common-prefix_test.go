package find_the_length_of_the_longest_common_prefix

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input1   []int
		input2   []int
		expected int
	}{
		{input1: []int{1, 10, 100}, input2: []int{1000}, expected: 3},
		{input1: []int{1, 2, 3}, input2: []int{4, 4, 4}, expected: 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := longestCommonPrefix(tc.input1, tc.input2)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
