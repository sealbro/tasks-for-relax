package path_sum

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	target := 22

	actual := hasPathSum(tree, target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	target := 5

	actual := hasPathSum(tree, target)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	var tree *TreeNode
	target := 0

	actual := hasPathSum(tree, target)

	assert.False(t, actual)
}

func TestCase4(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	target := 1

	actual := hasPathSum(tree, target)

	assert.False(t, actual)
}

func TestCase5(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
	}
	target := 1

	actual := hasPathSum(tree, target)

	assert.True(t, actual)
}

func TestCase6(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
	}
	target := 6

	actual := hasPathSum(tree, target)

	assert.False(t, actual)
}
