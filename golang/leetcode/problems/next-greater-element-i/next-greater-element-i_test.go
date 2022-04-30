package next_greater_element_i

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	num1 := []int{4, 1, 2}
	num2 := []int{1, 3, 4, 2}
	expected := []int{-1, 3, -1}

	actual := nextGreaterElement(num1, num2)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	num1 := []int{2, 4}
	num2 := []int{1, 2, 3, 4}
	expected := []int{3, -1}

	actual := nextGreaterElement(num1, num2)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	num1 := []int{5}
	num2 := []int{3}
	expected := []int{-1}

	actual := nextGreaterElement(num1, num2)

	assert.EqualMany(t, expected, actual)
}

func TestCase4(t *testing.T) {
	num1 := []int{7, 3, 5}
	num2 := []int{3, 4, 5, 1, 7}
	expected := []int{-1, 4, 7}

	actual := nextGreaterElement(num1, num2)

	assert.EqualMany(t, expected, actual)
}
