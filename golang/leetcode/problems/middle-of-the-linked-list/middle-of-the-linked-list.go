package middle_of_the_linked_list

// https://leetcode.com/problems/middle-of-the-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	current := head
	count := 0
	for current.Next != nil {
		current = current.Next
		count++
	}

	index := (count / 2) + (count % 2)
	current = head
	for i := 0; i < count+1; i++ {

		if i == index {
			return current
		}

		current = current.Next
	}

	return nil
}
