package reverse_string

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	actual := []byte{'h', 'e', 'l', 'l', 'o'}
	expected := []byte{'o', 'l', 'l', 'e', 'h'}

	reverseString(actual)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	actual := []byte{'h', 'i'}
	expected := []byte{'i', 'h'}

	reverseString(actual)

	assert.EqualMany(t, expected, actual)
}
