package sign_of_the_product_of_an_array

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	expected := 1

	actual := arraySign(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1, 5, 0, 2, -3}
	expected := 0

	actual := arraySign(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{-1, 1, -1, 1, -1}
	expected := -1

	actual := arraySign(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	expected := -1

	actual := arraySign(nums)

	assert.Equal(t, expected, actual)
}
