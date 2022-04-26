package path_sum

// https://leetcode.com/problems/path-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return detour(root, targetSum, 0, 0)
}

func detour(root *TreeNode, targetSum int, sum, level int) bool {
	if root == nil {
		return false
	}

	sum += root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == sum
	}

	return detour(root.Left, targetSum, sum, level+1) || detour(root.Right, targetSum, sum, level+1)
}
