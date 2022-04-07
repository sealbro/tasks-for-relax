package populating_next_right_pointers_in_each_node

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:   4,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   5,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val:   6,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   7,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Next: nil,
	}

	level1Right2 := &Node{
		Val:   7,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	level1Left2 := &Node{
		Val:   6,
		Left:  nil,
		Right: nil,
		Next:  level1Right2,
	}

	level1Right1 := &Node{
		Val:   5,
		Left:  nil,
		Right: nil,
		Next:  level1Left2,
	}
	level2Left1 := &Node{
		Val:   4,
		Left:  nil,
		Right: nil,
		Next:  level1Right1,
	}

	level1Right := &Node{
		Val:   3,
		Left:  level1Left2,
		Right: level1Right2,
		Next:  nil,
	}
	level1Left := &Node{
		Val:   2,
		Left:  level2Left1,
		Right: level1Right1,
		Next:  level1Right,
	}
	expected := &Node{
		Val:   1,
		Left:  level1Left,
		Right: level1Right,
		Next:  nil,
	}

	actual := connect(input)

	EqualTreeNode(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:   4,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Right: &Node{
			Val: 3,
			Right: &Node{
				Val:   7,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Next: nil,
	}

	level1Right2 := &Node{
		Val:   7,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	level2Left1 := &Node{
		Val:   4,
		Left:  nil,
		Right: nil,
		Next:  level1Right2,
	}

	level1Right := &Node{
		Val:   3,
		Left:  nil,
		Right: level1Right2,
		Next:  nil,
	}
	level1Left := &Node{
		Val:   2,
		Left:  level2Left1,
		Right: nil,
		Next:  level1Right,
	}
	expected := &Node{
		Val:   1,
		Left:  level1Left,
		Right: level1Right,
		Next:  nil,
	}

	actual := connect(input)

	EqualTreeNode(t, expected, actual)
}

func TestCase3(t *testing.T) {
	var input *Node
	var expected *Node

	actual := connect(input)

	EqualTreeNode(t, expected, actual)
}

func EqualTreeNode(t *testing.T, expected *Node, actual *Node) {
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

	expectedNext := expected
	actualNext := actual
	for {
		if expectedNext == nil && actualNext == nil {
			break
		}

		if expectedNext != nil && actualNext == nil || expectedNext == nil && actualNext != nil || expectedNext.Val != actualNext.Val {
			t.Errorf("\nExpected: %v,\nActual: %v.", expectedNext, actualNext)
			return
		}

		expectedNext = expectedNext.Next
		actualNext = actualNext.Next
	}

	EqualTreeNode(t, expected.Left, expected.Left)
	EqualTreeNode(t, expected.Right, expected.Right)
}
