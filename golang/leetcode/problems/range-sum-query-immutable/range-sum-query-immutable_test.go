package range_sum_query_immutable

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	ints := []int{-2, 0, 3, -5, 2, -1}
	arr := Constructor(ints)

	assert.Equal(t, 1, arr.SumRange(0, 2))
	assert.Equal(t, -1, arr.SumRange(2, 5))
	assert.Equal(t, -3, arr.SumRange(0, 5))
}
