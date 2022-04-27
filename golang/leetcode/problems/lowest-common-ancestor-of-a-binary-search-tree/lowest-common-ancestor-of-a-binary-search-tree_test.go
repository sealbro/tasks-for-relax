package lowest_common_ancestor_of_a_binary_search_tree

import (
	"testing"
)

func TestCase1(t *testing.T) {
	p := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
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
	q := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
	}
	tree := &TreeNode{
		Val:   6,
		Left:  p,
		Right: q,
	}

	expected := tree

	actual := lowestCommonAncestor(tree, p, q)

	EqualTreeNode(t, expected, actual)
}

func TestCase2(t *testing.T) {
	q := &TreeNode{
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
	}
	p := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: q,
	}
	tree := &TreeNode{
		Val:  6,
		Left: p,
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
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

	expected := p

	actual := lowestCommonAncestor(tree, p, q)

	EqualTreeNode(t, expected, actual)
}

func TestCase3(t *testing.T) {
	q := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	tree := &TreeNode{
		Val:   2,
		Left:  q,
		Right: nil,
	}

	p := tree

	expected := tree

	actual := lowestCommonAncestor(tree, p, q)

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
