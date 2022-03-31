package longest_substring_without_repeating_characters

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "abcabcbb"
	expected := 3

	actual := lengthOfLongestSubstring(input)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := "bbbbb"
	expected := 1

	actual := lengthOfLongestSubstring(input)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := "pwwkew"
	expected := 3

	actual := lengthOfLongestSubstring(input)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := "dvdf"
	expected := 3

	actual := lengthOfLongestSubstring(input)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := "ckilbkd"

	expected := 5

	actual := lengthOfLongestSubstring(input)

	assert.Equal(t, expected, actual)
}
