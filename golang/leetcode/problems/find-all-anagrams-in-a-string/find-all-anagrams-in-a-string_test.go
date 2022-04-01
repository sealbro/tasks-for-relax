package find_all_anagrams_in_a_string

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	s1 := "cbaebabacd"
	s2 := "abc"

	expected := []int{0, 6}

	actual := findAnagrams(s1, s2)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	s1 := "abab"
	s2 := "ab"

	expected := []int{0, 1, 2}

	actual := findAnagrams(s1, s2)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	s1 := "abacabacbac"
	s2 := "abc"

	expected := []int{1, 3, 5, 6, 7, 8}

	actual := findAnagrams(s1, s2)

	assert.EqualMany(t, expected, actual)
}
