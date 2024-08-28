package distribute_elements_into_two_arrays_i

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{2, 1, 3}, expected: []int{2, 3, 1}},
		{input: []int{5, 4, 3, 8}, expected: []int{5, 3, 4, 8}},
		{input: []int{6, 5, 14, 15}, expected: []int{6, 14, 15, 5}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := resultArray(tc.input)

			assert.EqualMany(t, tc.expected, actual)
		})
	}
}
