package reverse_bits

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	var target uint32 = 1
	var expected uint32 = 2147483648

	actual := reverseBits(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	var target uint32 = 2147483648
	var expected uint32 = 1

	actual := reverseBits(target)

	assert.Equal(t, expected, actual)
}
func TestCase3(t *testing.T) {
	var target uint32 = 3221225472
	var expected uint32 = 3

	actual := reverseBits(target)

	assert.Equal(t, expected, actual)
}
