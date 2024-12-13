package dot_product_of_two_sparse_vectors

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input1   SparseVector
		input2   SparseVector
		expected int
	}{
		{
			input1:   Constructor([]int{1, 0, 0, 2, 3}),
			input2:   Constructor([]int{0, 3, 0, 4, 0}),
			expected: 8,
		},
		{
			input1:   Constructor([]int{0, 1, 0, 0, 0}),
			input2:   Constructor([]int{0, 0, 0, 0, 2}),
			expected: 0,
		},
		{
			input1:   Constructor([]int{0, 1, 0, 0, 2, 0, 0}),
			input2:   Constructor([]int{1, 0, 0, 0, 3, 0, 4}),
			expected: 6,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase #%d", i+1), func(t *testing.T) {
			result := testCase.input1.dotProduct(testCase.input2)

			assert.Equal(t, testCase.expected, result)
		})
	}
}
