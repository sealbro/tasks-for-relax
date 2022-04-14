package power_of_two

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 1

	actual := isPowerOfTwo(target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	target := 16

	actual := isPowerOfTwo(target)

	assert.True(t, actual)
}

func TestCase3(t *testing.T) {
	target := 3

	actual := isPowerOfTwo(target)

	assert.False(t, actual)
}

func TestCase4(t *testing.T) {
	target := 7

	actual := isPowerOfTwo(target)

	assert.False(t, actual)
}

func TestCase5(t *testing.T) {
	target := -2147483648

	actual := isPowerOfTwo(target)

	assert.False(t, actual)
}
