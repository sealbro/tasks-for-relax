package symmetric_tree

// https://leetcode.com/problems/symmetric-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return detour(root.Left, root.Right)
}

func detour(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
		return false
	} else if node1.Val != node2.Val {
		return false
	}

	left := detour(node1.Left, node2.Right)

	if !left {
		return false
	}

	right := detour(node1.Right, node2.Left)

	return right
}
