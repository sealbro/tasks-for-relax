package sum_of_left_leaves

// https://leetcode.com/problems/sum-of-left-leaves/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sum
	}

	return sum
}
