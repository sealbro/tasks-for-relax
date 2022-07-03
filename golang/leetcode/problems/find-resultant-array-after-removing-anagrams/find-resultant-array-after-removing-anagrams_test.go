package find_resultant_array_after_removing_anagrams

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []string{"abba", "baba", "bbaa", "cd", "cd"}
	expected := []string{"abba", "cd"}

	actual := removeAnagrams(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e"}
	expected := []string{"a", "b", "c", "d", "e"}

	actual := removeAnagrams(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []string{"z", "z", "z", "gsw", "wsg", "gsw", "krptu"}
	expected := []string{"z", "gsw", "krptu"}

	actual := removeAnagrams(input)

	assert.EqualMany(t, expected, actual)
}
