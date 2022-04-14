package contains_duplicate

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 3, 1}

	actual := containsDuplicate(nums)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	actual := containsDuplicate(nums)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	actual := containsDuplicate(nums)

	assert.True(t, actual)
}
