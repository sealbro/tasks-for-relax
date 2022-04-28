package subtract_the_product_and_sum_of_digits_of_an_integer

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := 234
	expected := 15

	actual := subtractProductAndSum(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := 4421
	expected := 21

	actual := subtractProductAndSum(target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	target := 1
	expected := 0

	actual := subtractProductAndSum(target)

	assert.Equal(t, expected, actual)
}
