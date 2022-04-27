package two_sum_iv_input_is_a_bst

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   2,
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
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	target := 9

	actual := findTarget(nums, target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	nums := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   2,
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
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	target := 28

	actual := findTarget(nums, target)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	nums := &TreeNode{
		Val: 1,
	}
	target := 2

	actual := findTarget(nums, target)

	assert.False(t, actual)
}

func TestCase4(t *testing.T) {
	nums := &TreeNode{
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
	}
	target := 4

	actual := findTarget(nums, target)

	assert.True(t, actual)
}
