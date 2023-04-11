package defanging_an_ip_address

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	target := "1.1.1.1"
	expected := "1[.]1[.]1[.]1"

	actual := defangIPaddr(target)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	target := "255.100.50.0"
	expected := "255[.]100[.]50[.]0"

	actual := defangIPaddr(target)

	assert.Equal(t, expected, actual)
}
