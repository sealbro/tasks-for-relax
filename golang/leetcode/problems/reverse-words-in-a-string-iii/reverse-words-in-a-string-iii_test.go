package reverse_words_in_a_string_iii

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "Let's take LeetCode contest"
	expected := "s'teL ekat edoCteeL tsetnoc"

	actual := reverseWords(input)

	assert.Equal(t, expected, actual)
}
