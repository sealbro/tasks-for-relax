package divisor_game

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 2

	actual := divisorGame(target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	target := 3

	actual := divisorGame(target)

	assert.False(t, actual)
}
