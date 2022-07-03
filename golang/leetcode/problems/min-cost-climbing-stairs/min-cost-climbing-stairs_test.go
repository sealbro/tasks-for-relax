package min_cost_climbing_stairs

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{10, 15, 20}
	expected := 15

	actual := minCostClimbingStairs(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	expected := 6

	actual := minCostClimbingStairs(nums)

	assert.Equal(t, expected, actual)
}
