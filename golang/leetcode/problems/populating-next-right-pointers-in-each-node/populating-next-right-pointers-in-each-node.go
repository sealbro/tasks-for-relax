package populating_next_right_pointers_in_each_node

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	levels := make(map[int]*Node)

	connectBranches(root.Left, root.Right, levels, 1)

	return root
}

func connectBranches(left *Node, right *Node, levels map[int]*Node, level int) {
	levelLastNode := levels[level]
	if left != nil {
		if levelLastNode != nil {
			levelLastNode.Next = left
		}

		levels[level] = left
	}

	levelLastNode = levels[level]
	if right != nil {
		if levelLastNode != nil {
			levelLastNode.Next = right
		}
		levels[level] = right
	}

	if left != nil {
		connectBranches(left.Left, left.Right, levels, level+1)
	}

	if right != nil {
		connectBranches(right.Left, right.Right, levels, level+1)
	}
}
