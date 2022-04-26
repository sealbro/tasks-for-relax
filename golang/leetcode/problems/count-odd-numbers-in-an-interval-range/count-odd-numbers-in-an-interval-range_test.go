package count_odd_numbers_in_an_interval_range

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	low := 3
	high := 7
	expected := 3

	actual := countOdds(low, high)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	low := 8
	high := 10
	expected := 1

	actual := countOdds(low, high)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	low := 5
	high := 10
	expected := 3

	actual := countOdds(low, high)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	low := 6
	high := 9
	expected := 2

	actual := countOdds(low, high)

	assert.Equal(t, expected, actual)
}
