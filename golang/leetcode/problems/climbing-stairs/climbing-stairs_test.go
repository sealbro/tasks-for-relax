package climbing_stairs

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 2
	expected := 2

	actual := climbStairs(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := 3
	expected := 3

	actual := climbStairs(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	target := 45
	expected := 1836311903

	actual := climbStairs(target)

	assert.Equal(t, expected, actual)
}
