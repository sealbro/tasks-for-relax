package merge_two_sorted_lists

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	var result *ListNode
	var current *ListNode

	for list1 != nil || list2 != nil {
		if result == nil {
			result = &ListNode{}
			current = result
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}

		if (list1 != nil && list2 == nil) || (list1 != nil && list2 != nil && list1.Val < list2.Val) {
			current.Val = list1.Val
			list1 = list1.Next
		} else if list2 != nil {
			current.Val = list2.Val
			list2 = list2.Next
		}
	}

	return result
}
