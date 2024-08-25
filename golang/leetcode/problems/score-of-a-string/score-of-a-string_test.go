package score_of_a_string

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"hello", 13},
		{"zaz", 50},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := scoreOfString(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
