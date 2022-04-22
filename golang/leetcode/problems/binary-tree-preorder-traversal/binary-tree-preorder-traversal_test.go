package binary_tree_preorder_traversal

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	expected := []int{1, 2, 3}

	actual := preorderTraversal(tree)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	expected := []int{1}

	actual := preorderTraversal(tree)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	var tree *TreeNode
	var expected []int

	actual := preorderTraversal(tree)

	assert.EqualMany(t, expected, actual)
}
