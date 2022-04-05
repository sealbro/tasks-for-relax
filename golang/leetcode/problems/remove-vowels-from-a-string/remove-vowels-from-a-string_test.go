package remove_vowels_from_a_string

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "leetcodeisacommunityforcoders"
	expected := "ltcdscmmntyfrcdrs"

	actual := removeVowels(input)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := "aeiou"
	expected := ""

	actual := removeVowels(input)

	assert.Equal(t, expected, actual)
}
