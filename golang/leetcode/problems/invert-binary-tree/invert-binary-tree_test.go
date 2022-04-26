package invert_binary_tree

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	expected := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}

	actual := invertTree(input)

	EqualTreeNode(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	expected := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	actual := invertTree(input)

	EqualTreeNode(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	expected := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	actual := invertTree(input)

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
