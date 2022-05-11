package sqrtx

import (
	"golang/assert"
	"math"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 4
	expected := 2

	actual := mySqrt(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := 25
	expected := 5

	actual := mySqrt(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	target := 9
	expected := 3

	actual := mySqrt(target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	target := 0
	expected := 0

	actual := mySqrt(target)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	target := math.MaxUint32
	expected := 65535

	actual := mySqrt(target)

	assert.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	target := 8
	expected := 2

	actual := mySqrt(target)

	assert.Equal(t, expected, actual)
}

func TestCase7(t *testing.T) {
	target := 1
	expected := 1

	actual := mySqrt(target)

	assert.Equal(t, expected, actual)
}

func TestCase8(t *testing.T) {
	target := 2
	expected := 1

	actual := mySqrt(target)

	assert.Equal(t, expected, actual)
}
