package remove_duplicates_from_sorted_list

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head

	for current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
