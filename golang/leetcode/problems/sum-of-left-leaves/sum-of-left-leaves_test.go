package sum_of_left_leaves

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	root := &TreeNode{
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
	expected := 24

	actual := sumOfLeftLeaves(root)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	expected := 0

	actual := sumOfLeftLeaves(root)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   10,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val:   10,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	expected := 20

	actual := sumOfLeftLeaves(root)

	assert.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	expected := 4

	actual := sumOfLeftLeaves(root)

	assert.Equal(t, expected, actual)
}
