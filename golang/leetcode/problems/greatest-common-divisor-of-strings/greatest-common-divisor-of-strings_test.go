package greatest_common_divisor_of_strings

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	str1 := "ABCABC"
	str2 := "ABC"
	expected := "ABC"

	actual := gcdOfStrings(str1, str2)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	str1 := "ABABAB"
	str2 := "ABAB"
	expected := "AB"

	actual := gcdOfStrings(str1, str2)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	str1 := "LEET"
	str2 := "CODE"
	expected := ""

	actual := gcdOfStrings(str1, str2)

	assert.Equal(t, expected, actual)
}
