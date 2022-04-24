package symmetric_tree

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	actual := isSymmetric(tree)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	actual := isSymmetric(tree)

	assert.True(t, actual)
}

func TestCase3(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}

	actual := isSymmetric(tree)

	assert.False(t, actual)
}
