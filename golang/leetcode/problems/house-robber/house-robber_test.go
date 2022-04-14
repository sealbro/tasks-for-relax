package house_robber

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 4

	actual := rob(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	expected := 12

	actual := rob(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88, 43, 37, 157, 43, 221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80, 49, 245, 63, 25, 194, 72, 80, 143, 55, 209, 18, 55, 122, 65, 66, 177, 101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138, 212, 192, 125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83, 123, 120, 131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225, 54, 103, 55, 166, 124}
	expected := 7102

	actual := rob(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{1, 9, 3, 2, 9, 1}
	expected := 18

	actual := rob(nums)

	assert.Equal(t, expected, actual)
}
