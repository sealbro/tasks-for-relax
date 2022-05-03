package merge_strings_alternately

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	word1 := "abc"
	word2 := "pqr"
	expected := "apbqcr"

	actual := mergeAlternately(word1, word2)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"
	expected := "apbqrs"

	actual := mergeAlternately(word1, word2)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"
	expected := "apbqcd"

	actual := mergeAlternately(word1, word2)

	assert.Equal(t, expected, actual)
}
