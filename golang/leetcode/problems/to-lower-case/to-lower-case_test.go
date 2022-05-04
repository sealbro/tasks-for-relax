package to_lower_case

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "Hello"
	expected := "hello"

	actual := toLowerCase(input)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := "LOVELy"
	expected := "lovely"

	actual := toLowerCase(input)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := "s"
	expected := "s"

	actual := toLowerCase(input)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := "s48$^&*(DDD"
	expected := "s48$^&*(ddd"

	actual := toLowerCase(input)

	assert.Equal(t, expected, actual)
}
