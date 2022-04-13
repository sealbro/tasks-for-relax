package letter_case_permutation

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "a1b2"
	expected := []string{"a1b2", "A1b2", "a1B2", "A1B2"}

	actual := letterCasePermutation(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := "3z4"
	expected := []string{"3z4", "3Z4"}

	actual := letterCasePermutation(input)

	assert.EqualMany(t, expected, actual)
}
