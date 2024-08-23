package smallest_missing_integer_greater_than_sequential_prefix_sum

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 2, 5}, 6},
		{[]int{3, 4, 5, 1, 12, 14, 13}, 15},
		{[]int{4, 5, 6, 7, 8, 8, 9, 4, 3, 2, 7}, 30},
		{[]int{6, 4, 3}, 7},
		{[]int{10}, 11},
		{[]int{29, 30, 31, 32, 33, 34, 35, 36, 37}, 297},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := missingInteger(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
