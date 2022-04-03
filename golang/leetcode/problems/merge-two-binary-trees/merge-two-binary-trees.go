package merge_two_binary_trees

// https://leetcode.com/problems/merge-two-binary-trees/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
		return root1
	} else if root1 != nil && root2 == nil {
		return root1
	} else if root1 == nil && root2 != nil {
		return root2
	}
	return nil
}
