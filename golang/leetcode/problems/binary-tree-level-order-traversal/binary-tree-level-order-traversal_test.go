package binary_tree_level_order_traversal

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	input := &TreeNode{
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
	expected := [][]int{{3}, {9, 20}, {15, 7}}

	actual := levelOrder(input)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	expected := [][]int{{1}}

	actual := levelOrder(input)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase3(t *testing.T) {
	var input *TreeNode
	expected := [][]int{}

	actual := levelOrder(input)

	assert.EqualMultiDimensional(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  9,
			Left: nil,
			Right: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val:   10,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   11,
					Left:  nil,
					Right: nil,
				},
			},
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
	expected := [][]int{{3}, {9, 20}, {30, 15, 7}, {10, 11}}

	actual := levelOrder(input)

	assert.EqualMultiDimensional(t, expected, actual)
}
