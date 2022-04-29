package check_if_one_string_swap_can_make_strings_equal

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	s1 := "bank"
	s2 := "kanb"

	actual := areAlmostEqual(s1, s2)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	s1 := "attack"
	s2 := "defend"

	actual := areAlmostEqual(s1, s2)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	s1 := "abcd"
	s2 := "dcba"

	actual := areAlmostEqual(s1, s2)

	assert.False(t, actual)
}

func TestCase4(t *testing.T) {
	s1 := "bqzetqwrqegeupansshukcmnnezmooywwthvzkcciplknjfpzgpeybhuufoqnijzp"
	s2 := "gcsozkpencquoypwbhhzptnwerkqjmbmtzeokanefjifguniznwyuupqqhelzpscv"

	actual := areAlmostEqual(s1, s2)

	assert.False(t, actual)
}
