package kids_with_the_greatest_number_of_candies

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	expected := []bool{true, true, true, false, true}

	actual := kidsWithCandies(input, extraCandies)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	expected := []bool{true, false, false, false, false}

	actual := kidsWithCandies(input, extraCandies)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{12, 1, 12}
	extraCandies := 10
	expected := []bool{true, false, true}

	actual := kidsWithCandies(input, extraCandies)

	assert.EqualMany(t, expected, actual)
}
