package split_the_array

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 1, 2, 2, 3, 4}, true},
		{[]int{1, 1, 1, 1}, false},
		{[]int{6, 1, 3, 1, 1, 8, 9, 2}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := isPossibleToSplit(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
