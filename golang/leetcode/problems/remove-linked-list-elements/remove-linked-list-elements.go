package remove_linked_list_elements

// https://leetcode.com/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := firstNotEqualNode(head, val)

	current := newHead
	for current != nil {
		if current.Next == nil {
			break
		}

		if current.Next.Val == val {
			current.Next = firstNotEqualNode(current.Next.Next, val)
		}

		current = current.Next
	}

	return newHead
}

func firstNotEqualNode(head *ListNode, val int) *ListNode {
	newHead := head
	for newHead != nil {
		if newHead.Val == val {
			newHead = newHead.Next
		} else {
			return newHead
		}
	}

	return nil
}
