package goal_parser_interpretation

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	str := "G()(al)"
	expected := "Goal"

	actual := interpret(str)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	str := "G()()()()(al)"
	expected := "Gooooal"

	actual := interpret(str)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	str := "(al)G(al)()()G"
	expected := "alGalooG"

	actual := interpret(str)

	assert.Equal(t, expected, actual)
}
