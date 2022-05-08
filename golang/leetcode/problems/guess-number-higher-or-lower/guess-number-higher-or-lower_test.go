package guess_number_higher_or_lower

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 10
	expected := 6
	number = expected

	actual := guessNumber(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := 1
	expected := 1
	number = expected

	actual := guessNumber(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	target := 2
	expected := 1
	number = expected

	actual := guessNumber(target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	target := 2
	expected := 2
	number = expected

	actual := guessNumber(target)

	assert.Equal(t, expected, actual)
}
