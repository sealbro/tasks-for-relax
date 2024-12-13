package print_immutable_linked_list_in_reverse

import "fmt"

// https://leetcode.com/problems/print-immutable-linked-list-in-reverse/

type ImmutableListNode interface {
	getNext() ImmutableListNode
	printValue()
}

type ImmutableListNodeImp struct {
	value int
	next  *ImmutableListNodeImp
}

func (n *ImmutableListNodeImp) getNext() ImmutableListNode {
	if n == nil || n.next == nil {
		return nil
	}

	return n.next
}

func (n *ImmutableListNodeImp) printValue() {
	fmt.Println(n.value)
}

func printLinkedListInReverse(head ImmutableListNode) {
	if head == nil {
		return
	}

	printLinkedListInReverse(head.getNext())
	head.printValue()
}
