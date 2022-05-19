package sum_of_square_numbers

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 5

	actual := judgeSquareSum(target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	target := 3

	actual := judgeSquareSum(target)

	assert.False(t, actual)
}
