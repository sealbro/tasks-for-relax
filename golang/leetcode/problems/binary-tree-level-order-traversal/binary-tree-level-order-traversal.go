package binary_tree_level_order_traversal

// https://leetcode.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int

	detour(&result, root, 0)

	return result
}

func detour(result *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}

	for len(*result) < level+1 {
		*result = append(*result, []int{})
	}

	(*result)[level] = append((*result)[level], root.Val)

	detour(result, root.Left, level+1)
	detour(result, root.Right, level+1)
}
