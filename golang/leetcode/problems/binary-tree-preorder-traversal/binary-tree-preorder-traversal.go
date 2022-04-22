package binary_tree_preorder_traversal

// https://leetcode.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int

	insert(&result, root)

	return result
}

func insert(result *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	insert(result, root.Left)
	insert(result, root.Right)
}
