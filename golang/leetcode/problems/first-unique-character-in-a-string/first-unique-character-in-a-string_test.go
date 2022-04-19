package first_unique_character_in_a_string

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "leetcode"
	expected := 0

	actual := firstUniqChar(s)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	s := "loveleetcode"
	expected := 2

	actual := firstUniqChar(s)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	s := "aabb"
	expected := -1

	actual := firstUniqChar(s)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	s := "a"
	expected := 0

	actual := firstUniqChar(s)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	s := "aaaaaaaaaaaaaabbbbbbbbbbbbbc"
	expected := 27

	actual := firstUniqChar(s)

	assert.Equal(t, expected, actual)
}
