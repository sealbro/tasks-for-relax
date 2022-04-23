package binary_tree_inorder_traversal

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := &TreeNode{
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
	expected := []int{1, 3, 2}

	actual := inorderTraversal(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	expected := []int{1}

	actual := inorderTraversal(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	var input *TreeNode
	expected := []int{}

	actual := inorderTraversal(input)

	assert.EqualMany(t, expected, actual)
}
