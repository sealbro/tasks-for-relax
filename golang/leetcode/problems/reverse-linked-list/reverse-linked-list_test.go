package reverse_linked_list

import (
	"testing"
)

func TestCase1(t *testing.T) {
	var list = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	var expected = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	actual := reverseList(list)

	EqualLinkedList(t, actual, expected)
}

func TestCase2(t *testing.T) {
	var list = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	var expected = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	actual := reverseList(list)

	EqualLinkedList(t, actual, expected)
}

func TestCase3(t *testing.T) {
	var list *ListNode

	var expected *ListNode

	actual := reverseList(list)

	EqualLinkedList(t, actual, expected)
}

func EqualLinkedList(t *testing.T, actual *ListNode, expected *ListNode) {
	for actual != nil && expected != nil {
		if actual.Val != expected.Val {
			t.Errorf("Expected: %v, Actual: %v.", expected, actual)
			break
		}

		actual = actual.Next
		expected = expected.Next
	}
}
