package average_salary_excluding_the_minimum_and_maximum_salary

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{4000, 3000, 1000, 2000}
	expected := 2500.00000

	actual := average(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{1000, 2000, 3000}
	expected := 2000.00000

	actual := average(nums)

	assert.Equal(t, expected, actual)
}
