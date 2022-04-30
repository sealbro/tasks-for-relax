package n_ary_tree_preorder_traversal

import (
	"golang/assert"
	"testing"
)

//1,null,3,2,4,null,5,6

func TestCase1(t *testing.T) {
	input := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	expected := []int{1, 3, 5, 6, 2, 4}

	actual := preorder(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	var input *Node = nil
	var expected []int

	actual := preorder(input)

	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	var input = &Node{
		Val:      1,
		Children: nil,
	}
	expected := []int{1}

	actual := preorder(input)

	assert.EqualMany(t, expected, actual)
}
