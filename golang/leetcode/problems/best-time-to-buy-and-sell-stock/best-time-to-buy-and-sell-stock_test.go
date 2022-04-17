package best_time_to_buy_and_sell_stock

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	expected := 5

	actual := maxProfit(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{7, 6, 4, 3, 1}
	expected := 0

	actual := maxProfit(nums)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4, 10}
	expected := 9

	actual := maxProfit(nums)

	assert.Equal(t, expected, actual)
}
