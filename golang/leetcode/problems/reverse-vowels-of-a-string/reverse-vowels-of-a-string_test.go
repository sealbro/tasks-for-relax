package reverse_vowels_of_a_string

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := "hello"
	expected := "holle"

	actual := reverseVowels(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := "leetcode"
	expected := "leotcede"

	actual := reverseVowels(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	target := "aA"
	expected := "Aa"

	actual := reverseVowels(target)

	assert.Equal(t, expected, actual)
}
