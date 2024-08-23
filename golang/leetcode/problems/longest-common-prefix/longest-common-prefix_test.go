package longest_common_prefix

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"flower", "aflow", "rflight"}, ""},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := longestCommonPrefix(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
