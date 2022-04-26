package insert_into_a_binary_search_tree

import (
	"testing"
)

func TestCase1(t *testing.T) {
	nums := &TreeNode{
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
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}
	target := 5
	expected := &TreeNode{
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
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	actual := insertIntoBST(nums, target)

	EqualTreeNode(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := &TreeNode{
		Val: 40,
		Left: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   30,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 60,
			Left: &TreeNode{
				Val:   50,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   70,
				Left:  nil,
				Right: nil,
			},
		},
	}
	target := 25
	var expected = &TreeNode{
		Val: 40,
		Left: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   25,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 60,
			Left: &TreeNode{
				Val:   50,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   70,
				Left:  nil,
				Right: nil,
			},
		},
	}

	actual := insertIntoBST(nums, target)

	EqualTreeNode(t, expected, actual)
}

func TestCase3(t *testing.T) {
	var nums *TreeNode
	target := 25
	var expected = &TreeNode{
		Val: 25,
	}

	actual := insertIntoBST(nums, target)

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
