package convert_sorted_array_to_binary_search_tree

import (
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5}
	expected := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 0,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}

	actual := sortedArrayToBST(nums)

	EqualTreeNode(t, expected, actual)
}

func EqualTreeNode(t *testing.T, expected *TreeNode, actual *TreeNode) {
	if expected == nil && actual == nil {
		return
	}

	if expected != nil && actual == nil || expected == nil && actual != nil {
		t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
		return
	}

	if expected.Val != actual.Val {
		t.Errorf("\nExpected: %v,\nActual: %v.", expected.Val, actual.Val)
		return
	}

	EqualTreeNode(t, expected.Left, actual.Left)
	EqualTreeNode(t, expected.Right, actual.Right)
}
