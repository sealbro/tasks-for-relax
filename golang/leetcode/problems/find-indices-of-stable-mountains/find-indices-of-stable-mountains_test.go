package find_indices_of_stable_mountains

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		inputArr       []int
		inputThreshold int
		expected       []int
	}{
		{
			inputArr:       []int{1, 2, 3, 4, 5},
			inputThreshold: 2,
			expected:       []int{3, 4},
		},
		{
			inputArr:       []int{10, 1, 10, 1, 10},
			inputThreshold: 3,
			expected:       []int{1, 3},
		},
		{
			inputArr:       []int{10, 1, 10, 1, 10},
			inputThreshold: 10,
			expected:       []int{},
		},
		{
			inputArr:       []int{5, 10, 11, 12, 5, 10, 10},
			inputThreshold: 5,
			expected:       []int{2, 3, 4, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := stableMountains(tc.inputArr, tc.inputThreshold)

			assert.EqualMany(t, tc.expected, actual)
		})
	}
}
