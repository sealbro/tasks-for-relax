package binary_tree_inorder_traversal

// https://leetcode.com/problems/binary-tree-inorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int

	insert(&result, root)

	return result
}

func insert(result *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	insert(result, root.Left)
	*result = append(*result, root.Val)
	insert(result, root.Right)
}
