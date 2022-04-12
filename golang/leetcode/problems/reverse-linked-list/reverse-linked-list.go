package reverse_linked_list

// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head.Next
	head.Next = nil
	last := head
	for current != nil {
		temp := current.Next

		current.Next, last = last, current

		current = temp
	}

	return last
}
