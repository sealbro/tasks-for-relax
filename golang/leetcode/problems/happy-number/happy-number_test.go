package happy_number

import (
	"golang/assert"
	"math"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 19

	actual := isHappy(target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	target := 2

	actual := isHappy(target)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	target := math.MaxInt32

	actual := isHappy(target)

	assert.False(t, actual)
}

func TestCase4(t *testing.T) {
	target := 1

	actual := isHappy(target)

	assert.True(t, actual)
}

func TestCase5(t *testing.T) {
	target := 7

	actual := isHappy(target)

	assert.True(t, actual)
}
