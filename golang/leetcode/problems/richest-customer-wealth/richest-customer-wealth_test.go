package richest_customer_wealth

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := [][]int{{1, 2, 3}, {3, 2, 1}}
	expected := 6

	actual := maximumWealth(nums)

	assert.Equal(t, expected, actual)
}
func TestCase2(t *testing.T) {
	nums := [][]int{{1, 5}, {7, 3}, {3, 5}}
	expected := 10

	actual := maximumWealth(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	expected := 17

	actual := maximumWealth(nums)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := [][]int{{1}, {40}, {4}}
	expected := 40

	actual := maximumWealth(nums)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums := [][]int{{100}}
	expected := 100

	actual := maximumWealth(nums)

	assert.Equal(t, expected, actual)
}
