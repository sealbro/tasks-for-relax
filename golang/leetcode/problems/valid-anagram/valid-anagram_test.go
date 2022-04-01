package valid_anagram

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"

	actual := isAnagram(s1, s2)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	s1 := "rat"
	s2 := "car"

	actual := isAnagram(s1, s2)

	assert.False(t, actual)
}
