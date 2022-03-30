package remove_nth_node_from_end_of_list

import (
	"testing"
)

func TestCase1(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	n := 1

	actual := removeNthFromEnd(head, n)
	if actual != nil {
		t.Errorf("Expected: %v, Actual: %v.", nil, actual)
	}
}

func TestCase2(t *testing.T) {
	head := &ListNode{
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
	n := 2

	actual := removeNthFromEnd(head, n)

	expected := []int{1, 2, 3, 5}
	i := 0
	current := actual
	for current.Next != nil {
		if current.Val != expected[i] {
			t.Errorf("Expected: %v, Actual: %v.", expected, actual)
			break
		}
		current = current.Next
		i++
	}
}

func TestCase3(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	n := 2

	actual := removeNthFromEnd(head, n)

	expected := []int{2}
	i := 0
	current := actual
	for current.Next != nil {
		if current.Val != expected[i] {
			t.Errorf("Expected: %v, Actual: %v.", expected, actual.Val)
			break
		}
		current = current.Next
		i++
	}
}
