package fibonacci_number

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 2
	expected := 1

	actual := fib(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := 3
	expected := 2

	actual := fib(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	target := 4
	expected := 3

	actual := fib(target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	target := 7
	expected := 13

	actual := fib(target)

	assert.Equal(t, expected, actual)
}
