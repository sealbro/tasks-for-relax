package alice_and_bob_playing_flower_game

import (
	"fmt"
	"golang/assert"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		input1, input2 int
		expected       int64
	}{
		{input1: 3, input2: 2, expected: 3},
		{input1: 1, input2: 1, expected: 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := flowerGame(tc.input1, tc.input2)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
