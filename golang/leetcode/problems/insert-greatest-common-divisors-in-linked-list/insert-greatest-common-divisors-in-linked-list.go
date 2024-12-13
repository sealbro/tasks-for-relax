package insert_greatest_common_divisors_in_linked_list

// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	a := head.Val
	b := head.Next.Val
	for b != 0 {
		a, b = b, a%b
	}

	head.Next = &ListNode{
		Val:  a,
		Next: insertGreatestCommonDivisors(head.Next),
	}

	return head
}
