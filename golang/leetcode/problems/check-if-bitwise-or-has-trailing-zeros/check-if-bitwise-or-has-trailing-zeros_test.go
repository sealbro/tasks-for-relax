package check_if_bitwise_or_has_trailing_zeros

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
		{input: []int{1, 2, 3, 4, 5}, expected: true},
		{input: []int{2, 4, 8, 16}, expected: true},
		{input: []int{1, 3, 5, 7, 9}, expected: false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := hasTrailingZeros(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
