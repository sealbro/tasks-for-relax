package ransom_note

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	note := "a"
	magazine := "b"

	actual := canConstruct(note, magazine)

	assert.False(t, actual)
}

func TestCase2(t *testing.T) {
	note := "aa"
	magazine := "ab"

	actual := canConstruct(note, magazine)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	note := "aa"
	magazine := "aab"

	actual := canConstruct(note, magazine)

	assert.True(t, actual)
}
