package middle_of_the_linked_list

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: expected,
		},
	}

	actual := middleNode(head)

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v.", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: expected,
			},
		},
	}

	actual := middleNode(head)

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v.", expected, actual)
	}
}
