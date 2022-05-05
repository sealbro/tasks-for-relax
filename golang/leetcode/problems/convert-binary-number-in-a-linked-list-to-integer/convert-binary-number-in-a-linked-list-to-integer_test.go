package convert_binary_number_in_a_linked_list_to_integer

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	expected := 5

	actual := getDecimalValue(nums)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := &ListNode{
		Val: 0,
	}

	expected := 0

	actual := getDecimalValue(nums)

	assert.Equal(t, expected, actual)
}
