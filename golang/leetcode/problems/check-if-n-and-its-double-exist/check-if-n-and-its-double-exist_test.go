package check_if_n_and_its_double_exist

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{10, 2, 5, 3}

	actual := checkIfExist(nums)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{7, 1, 14, 11}

	actual := checkIfExist(nums)

	assert.True(t, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{3, 1, 7, 11}

	actual := checkIfExist(nums)

	assert.False(t, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{3, 6}

	actual := checkIfExist(nums)

	assert.True(t, actual)
}

func TestCase5(t *testing.T) {
	nums := []int{33, 9}

	actual := checkIfExist(nums)

	assert.False(t, actual)
}
