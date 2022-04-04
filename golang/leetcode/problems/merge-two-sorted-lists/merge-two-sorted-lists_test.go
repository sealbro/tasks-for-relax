package merge_two_sorted_lists

import (
	"testing"
)

func TestCase1(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}

	actual := mergeTwoLists(list1, list2)

	EqualLinkedList(t, actual, expected)
}

func TestCase2(t *testing.T) {
	var list1 *ListNode
	var list2 *ListNode
	var expected *ListNode

	actual := mergeTwoLists(list1, list2)

	EqualLinkedList(t, actual, expected)
}

func TestCase3(t *testing.T) {
	var list1 *ListNode
	var list2 = &ListNode{
		Val:  1,
		Next: nil,
	}
	var expected = &ListNode{
		Val:  1,
		Next: nil,
	}

	actual := mergeTwoLists(list1, list2)

	EqualLinkedList(t, actual, expected)
}

func TestCase4(t *testing.T) {
	var list1 = &ListNode{
		Val:  2,
		Next: nil,
	}
	var list2 = &ListNode{
		Val:  1,
		Next: nil,
	}
	var expected = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	actual := mergeTwoLists(list1, list2)

	EqualLinkedList(t, actual, expected)
}

func EqualLinkedList(t *testing.T, actual *ListNode, expected *ListNode) {
	for actual != nil && expected != nil {

		if actual.Val != expected.Val {
			t.Errorf("Expected: %v, Actual: %v.", expected, actual)
		}

		actual = actual.Next
		expected = expected.Next
	}
}
