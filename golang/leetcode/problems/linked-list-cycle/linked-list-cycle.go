package linked_list_cycle

import "fmt"

// https://leetcode.com/problems/linked-list-cycle/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next
	for slow != fast {
		if slow == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func hasCycleHash(head *ListNode) bool {
	current := head
	hashTable := make(map[string]struct{})
	for current != nil {
		address := fmt.Sprintf("%p", current)
		if _, ok := hashTable[address]; ok {
			return true
		} else {
			hashTable[address] = struct{}{}
		}
		current = current.Next
	}

	return false
}
