package burst_balloons

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	expected := 167

	actual := maxCoins(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1, 5}
	expected := 10

	actual := maxCoins(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{3, 7, 2, 9, 1, 5, 8}
	expected := 1235

	actual := maxCoins(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{15}
	expected := 15

	actual := maxCoins(nums)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums := []int{9, 76, 64, 21, 97}
	expected := 669494

	actual := maxCoins(nums)

	assert.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	nums := []int{9, 76, 64, 21, 97, 60}
	expected := 1086136

	actual := maxCoins(nums)

	assert.Equal(t, expected, actual)
}
