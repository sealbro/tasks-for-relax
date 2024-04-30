package roman_to_integer

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := "III"
	expected := 3

	actual := romanToInt(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := "LVIII"
	expected := 58

	actual := romanToInt(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	target := "MCMXCIV"
	expected := 1994

	actual := romanToInt(target)

	assert.Equal(t, expected, actual)
}
