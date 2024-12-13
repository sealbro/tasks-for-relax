package insert_greatest_common_divisors_in_linked_list

import (
	"fmt"
	"golang/assert"
	"testing"
)

func toListNode(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	head := &ListNode{
		Val: input[0],
	}

	current := head
	for i := 1; i < len(input); i++ {
		current.Next = &ListNode{
			Val: input[i],
		}
		current = current.Next
	}

	return head
}

func TestCase(t *testing.T) {
	testCases := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{
			input:    toListNode([]int{18, 6, 10, 3}),
			expected: toListNode([]int{18, 6, 6, 2, 10, 1, 3}),
		},
		{
			input:    toListNode([]int{7}),
			expected: toListNode([]int{7}),
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case #%d", i+1), func(t *testing.T) {
			result := insertGreatestCommonDivisors(testCase.input)

			head := testCase.expected
			for head != nil {
				assert.Equal(t, head.Val, result.Val)
				head = head.Next
				result = result.Next
			}
		})
	}

}
