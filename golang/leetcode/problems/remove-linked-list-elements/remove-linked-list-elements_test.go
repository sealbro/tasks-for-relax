package remove_linked_list_elements

import (
	"testing"
)

func TestCase1(t *testing.T) {
	nums := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	target := 6
	expected := &ListNode{
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

	actual := removeElements(nums, target)

	EqualLinkedList(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	target := 7
	var expected *ListNode = nil

	actual := removeElements(nums, target)

	EqualLinkedList(t, expected, actual)
}

func TestCase3(t *testing.T) {
	var nums *ListNode = nil
	target := 7
	var expected *ListNode = nil

	actual := removeElements(nums, target)

	EqualLinkedList(t, expected, actual)
}

func TestCase4(t *testing.T) {
	nums := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	target := 2
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	actual := removeElements(nums, target)

	EqualLinkedList(t, expected, actual)
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
