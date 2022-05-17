package arranging_coins

import (
	"golang/assert"
	"math"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 5
	expected := 2

	actual := arrangeCoins(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := 8
	expected := 3

	actual := arrangeCoins(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	target := 2
	expected := 1

	actual := arrangeCoins(target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	target := 1
	expected := 1

	actual := arrangeCoins(target)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	target := math.MaxInt32
	expected := 65535

	actual := arrangeCoins(target)

	assert.Equal(t, expected, actual)
}
