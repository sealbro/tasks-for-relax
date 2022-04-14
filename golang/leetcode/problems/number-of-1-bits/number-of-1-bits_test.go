package number_of_1_bits

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	var target uint32 = 11
	expected := 3

	actual := hammingWeight(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	var target uint32 = 3
	expected := 2

	actual := hammingWeight(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	var target uint32 = 0
	expected := 0

	actual := hammingWeight(target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	var target uint32 = 15
	expected := 4

	actual := hammingWeight(target)

	assert.Equal(t, expected, actual)
}
