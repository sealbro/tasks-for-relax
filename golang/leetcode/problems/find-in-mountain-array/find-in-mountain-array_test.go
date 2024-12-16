package find_in_mountain_array

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		target   int
		input    *MountainArray
		expected int
	}{
		{
			target:   3,
			input:    &MountainArray{arr: []int{1, 2, 3, 4, 5, 3, 1}},
			expected: 2,
		},
		{
			target:   3,
			input:    &MountainArray{arr: []int{0, 1, 2, 4, 2, 1}},
			expected: -1,
		},
		{
			target:   1,
			input:    &MountainArray{arr: []int{1, 5, 2}},
			expected: 0,
		},
		{
			target:   2,
			input:    &MountainArray{arr: []int{1, 5, 2}},
			expected: 2,
		},
		{
			target:   1,
			input:    &MountainArray{arr: []int{0, 5, 3, 1}},
			expected: 3,
		},
		{
			target:   3,
			input:    &MountainArray{arr: []int{3, 5, 3, 2, 0}},
			expected: 0,
		},
		{
			target:   3,
			input:    &MountainArray{arr: []int{1, 2, 3, 5, 3}},
			expected: 2,
		},
		{
			target:   22,
			input:    &MountainArray{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}},
			expected: -1,
		},
		{
			target:   81,
			input:    &MountainArray{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82}},
			expected: 80,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := findInMountainArray(tc.target, tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
