package valid_perfect_square

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 16

	actual := isPerfectSquare(target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	target := 14

	actual := isPerfectSquare(target)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	target := 25

	actual := isPerfectSquare(target)

	assert.True(t, actual)
}

func TestCase4(t *testing.T) {
	target := 1

	actual := isPerfectSquare(target)

	assert.True(t, actual)
}

func TestCase5(t *testing.T) {
	target := 9

	actual := isPerfectSquare(target)

	assert.True(t, actual)
}
