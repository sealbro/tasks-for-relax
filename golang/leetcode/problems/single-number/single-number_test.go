package single_number

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 2, 1}
	expected := 1

	actual := singleNumber(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	expected := 4

	actual := singleNumber(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2, 4, 20}
	expected := 20

	actual := singleNumber(nums)

	assert.Equal(t, expected, actual)
}
