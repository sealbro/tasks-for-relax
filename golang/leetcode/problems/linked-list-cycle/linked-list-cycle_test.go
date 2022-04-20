package linked_list_cycle

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	next := &ListNode{
		Val: 2,
	}
	next.Next = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  -4,
			Next: next,
		},
	}

	input := &ListNode{
		Val:  3,
		Next: next,
	}

	actual := hasCycle(input)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	input := &ListNode{
		Val: 1,
	}
	input.Next = &ListNode{
		Val:  2,
		Next: input,
	}

	actual := hasCycle(input)

	assert.True(t, actual)
}

func TestCase1h(t *testing.T) {
	next := &ListNode{
		Val: 2,
	}
	next.Next = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  -4,
			Next: next,
		},
	}

	input := &ListNode{
		Val:  3,
		Next: next,
	}

	actual := hasCycleHash(input)

	assert.True(t, actual)
}

func TestCase2h(t *testing.T) {
	input := &ListNode{
		Val: 1,
	}
	input.Next = &ListNode{
		Val:  2,
		Next: input,
	}

	actual := hasCycleHash(input)

	assert.True(t, actual)
}
