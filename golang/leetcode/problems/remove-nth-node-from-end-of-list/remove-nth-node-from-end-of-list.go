package remove_nth_node_from_end_of_list

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	count := 1
	for current.Next != nil {
		current = current.Next
		count++
	}

	if count == 1 && n == 1 {
		return nil
	}

	indexToRemove := count - n

	if indexToRemove == 0 {
		return head.Next
	}

	current = head
	prev := head
	for i := 0; i < count; i++ {
		if i == indexToRemove {
			prev.Next = current.Next
			break
		}

		prev = current
		current = current.Next
	}

	return head
}
