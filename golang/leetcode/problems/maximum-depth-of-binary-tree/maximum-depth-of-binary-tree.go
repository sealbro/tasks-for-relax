package maximum_depth_of_binary_tree

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return detour(root, 0)
}

func detour(root *TreeNode, level int) int {
	if root == nil {
		return level
	}

	left := detour(root.Left, level+1)
	right := detour(root.Right, level+1)

	if right < left {
		return left
	}
	return right
}
