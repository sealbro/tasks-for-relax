package company_cases_1

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	x, y, z := Case1()

	expected := []int{0, 1, 2, 4}

	assert.EqualMany(t, expected[:3], x)
	assert.EqualMany(t, expected, y)
	assert.EqualMany(t, expected, z)
}

func TestCase2(t *testing.T) {
	seconds := Case2()

	assert.Equal(t, 2, seconds)
}

func TestCase3(t *testing.T) {
	withoutZeros := Case3()

	assert.EqualMany(t, []int{2, 1, 1}, withoutZeros)
}

func TestCase4(t *testing.T) {
	count := Case4()

	assert.Equal(t, 100, count)
}
