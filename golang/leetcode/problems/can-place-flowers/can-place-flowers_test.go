package can_place_flowers

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 0, 0, 0, 1}
	target := 1

	actual := canPlaceFlowers(nums, target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1, 0, 0, 0, 1}
	target := 2

	actual := canPlaceFlowers(nums, target)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{0, 0, 1, 0, 1}
	target := 1

	actual := canPlaceFlowers(nums, target)

	assert.True(t, actual)
}

func TestCase4(t *testing.T) {
	nums := []int{1, 0, 1, 0, 0}
	target := 1

	actual := canPlaceFlowers(nums, target)

	assert.True(t, actual)
}

func TestCase5(t *testing.T) {
	nums := []int{1, 0, 0, 0, 0}
	target := 2

	actual := canPlaceFlowers(nums, target)

	assert.True(t, actual)
}

func TestCase6(t *testing.T) {
	nums := []int{0}
	target := 1

	actual := canPlaceFlowers(nums, target)

	assert.True(t, actual)
}

func TestCase7(t *testing.T) {
	nums := []int{0, 0, 1, 0, 0}
	target := 1

	actual := canPlaceFlowers(nums, target)

	assert.True(t, actual)
}
