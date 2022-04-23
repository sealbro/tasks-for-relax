package binary_tree_postorder_traversal

// https://leetcode.com/problems/binary-tree-postorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	insert(&result, root)
	return result
}

func insert(result *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	insert(result, root.Left)
	insert(result, root.Right)

	*result = append(*result, root.Val)
}
