package can_make_arithmetic_progression_from_sequence

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{3, 5, 1}

	actual := canMakeArithmeticProgression(nums)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 4}

	actual := canMakeArithmeticProgression(nums)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{1, 5, 3, -1, -3, -5}

	actual := canMakeArithmeticProgression(nums)

	assert.True(t, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{1, 2}

	actual := canMakeArithmeticProgression(nums)

	assert.True(t, actual)
}
