package convert_binary_number_in_a_linked_list_to_integer

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	value := head.Val
	current := head.Next
	for current != nil {
		value = (value << 1) | current.Val
		current = current.Next
	}

	return value
}
