package find_smallest_letter_greater_than_target

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []byte{'c', 'f', 'j'}
	var target byte = 'a'
	var expected byte = 'c'

	actual := nextGreatestLetter(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []byte{'c', 'f', 'j'}
	var target byte = 'c'
	var expected byte = 'f'

	actual := nextGreatestLetter(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []byte{'c', 'f', 'j'}
	var target byte = 'd'
	var expected byte = 'f'

	actual := nextGreatestLetter(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := []byte{'c', 'c', 'c', 'c', 'j'}
	var target byte = 'd'
	var expected byte = 'j'

	actual := nextGreatestLetter(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	nums := []byte{'c', 'j'}
	var target byte = 'd'
	var expected byte = 'j'

	actual := nextGreatestLetter(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	nums := []byte{'c', 'f', 'j'}
	var target byte = 'j'
	var expected byte = 'c'

	actual := nextGreatestLetter(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase7(t *testing.T) {
	nums := []byte{'a', 'b', 'd'}
	var target byte = 'c'
	var expected byte = 'd'

	actual := nextGreatestLetter(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCase8(t *testing.T) {
	nums := []byte{'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n', 'n', 'n'}
	var target byte = 'e'
	var expected byte = 'n'

	actual := nextGreatestLetter(nums, target)

	assert.Equal(t, expected, actual)
}
