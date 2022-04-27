package lowest_common_ancestor_of_a_binary_search_tree

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if (root.Val >= p.Val && root.Val <= q.Val) ||
		(root.Val <= p.Val && root.Val >= q.Val) {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return lowestCommonAncestor(root.Right, p, q)
}
