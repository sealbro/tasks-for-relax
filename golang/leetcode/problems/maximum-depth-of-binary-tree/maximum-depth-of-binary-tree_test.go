package maximum_depth_of_binary_tree

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	expected := 3

	actual := maxDepth(tree)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
		},
	}
	expected := 2

	actual := maxDepth(tree)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
	}
	expected := 1

	actual := maxDepth(tree)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	expected := 2

	actual := maxDepth(tree)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	var tree *TreeNode
	expected := 0

	actual := maxDepth(tree)

	assert.Equal(t, expected, actual)
}
