package remove_duplicates_from_sorted_list

import (
	"testing"
)

func TestCase1(t *testing.T) {
	nums := &ListNode{
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
							Val: 5,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

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

	actual := deleteDuplicates(nums)

	EqualLinkedList(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	expected := &ListNode{
		Val:  1,
		Next: nil,
	}

	actual := deleteDuplicates(nums)

	EqualLinkedList(t, expected, actual)
}

func TestCase3(t *testing.T) {
	var nums *ListNode = nil
	var expected *ListNode = nil

	actual := deleteDuplicates(nums)

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
