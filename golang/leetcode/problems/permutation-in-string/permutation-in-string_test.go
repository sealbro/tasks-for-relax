package permutation_in_string

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"

	actual := checkInclusion(s1, s2)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	s1 := "ab"
	s2 := "eidboaoo"

	actual := checkInclusion(s1, s2)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	s1 := "abc"
	s2 := "ccccbbbbaaaa"

	actual := checkInclusion(s1, s2)

	assert.False(t, actual)
}
