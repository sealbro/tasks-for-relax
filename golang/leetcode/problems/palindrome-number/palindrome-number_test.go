package palindrome_number

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 121

	actual := isPalindrome(target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	target := 123

	actual := isPalindrome(target)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	target := 0

	actual := isPalindrome(target)

	assert.True(t, actual)
}

func TestCase4(t *testing.T) {
	target := 10

	actual := isPalindrome(target)

	assert.False(t, actual)
}

func TestCase5(t *testing.T) {
	target := 11

	actual := isPalindrome(target)

	assert.True(t, actual)
}

func TestCase6(t *testing.T) {
	target := -11

	actual := isPalindrome(target)

	assert.False(t, actual)
}
