package valid_parentheses

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "()"

	actual := isValid(s)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	s := "()[]{}"

	actual := isValid(s)

	assert.True(t, actual)
}

func TestCase3(t *testing.T) {
	s := "(}"

	actual := isValid(s)

	assert.False(t, actual)
}

func TestCase4(t *testing.T) {
	s := "("

	actual := isValid(s)

	assert.False(t, actual)
}
