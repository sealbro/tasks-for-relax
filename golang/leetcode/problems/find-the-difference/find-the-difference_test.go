package find_the_difference

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	str1 := "abcd"
	str2 := "abcde"
	var expected byte = 'e'

	actual := findTheDifference(str1, str2)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	str1 := ""
	str2 := "y"
	var expected byte = 'y'

	actual := findTheDifference(str1, str2)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	str1 := "a"
	str2 := "za"
	var expected byte = 'z'

	actual := findTheDifference(str1, str2)

	assert.Equal(t, expected, actual)
}
