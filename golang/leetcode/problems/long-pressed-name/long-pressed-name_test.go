package long_pressed_name

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	str1 := "alex"
	str2 := "aaleex"

	actual := isLongPressedName(str1, str2)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	str1 := "saeed"
	str2 := "ssaaedd"

	actual := isLongPressedName(str1, str2)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	str1 := "saeed"
	str2 := "saaeeeed"

	actual := isLongPressedName(str1, str2)

	assert.True(t, actual)
}

func TestCase4(t *testing.T) {
	str1 := "saeed"
	str2 := "saaeeee"

	actual := isLongPressedName(str1, str2)

	assert.False(t, actual)
}

func TestCase5(t *testing.T) {
	str1 := "saeed"
	str2 := "saaeqqeeed"

	actual := isLongPressedName(str1, str2)

	assert.False(t, actual)
}

func TestCase6(t *testing.T) {
	str1 := "vtkgn"
	str2 := "vttkgnn"

	actual := isLongPressedName(str1, str2)

	assert.True(t, actual)
}

func TestCase7(t *testing.T) {
	str1 := "vtkgn"
	str2 := "vttkgnnq"

	actual := isLongPressedName(str1, str2)

	assert.False(t, actual)
}
