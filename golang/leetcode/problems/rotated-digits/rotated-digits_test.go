package rotated_digits

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 10
	expected := 4

	actual := rotatedDigits(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	target := 2
	expected := 1

	actual := rotatedDigits(target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	target := 857
	expected := 247

	actual := rotatedDigits(target)

	assert.Equal(t, expected, actual)
}
