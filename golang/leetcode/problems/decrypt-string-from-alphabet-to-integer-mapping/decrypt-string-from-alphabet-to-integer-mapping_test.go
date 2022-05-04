package decrypt_string_from_alphabet_to_integer_mapping

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "10#11#12"
	expected := "jkab"

	actual := freqAlphabets(s)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	s := "1326#"
	expected := "acz"

	actual := freqAlphabets(s)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	s := "10#1326#"
	expected := "jacz"

	actual := freqAlphabets(s)

	assert.Equal(t, expected, actual)
}
