package validate_binary_search_tree

// https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LastValue struct {
	init  bool
	value int
}

func isValidBST(root *TreeNode) bool {
	value := &LastValue{}
	return detour(root, value)
}

func detour(root *TreeNode, lastValue *LastValue) bool {
	if root == nil {
		return true
	}

	if !detour(root.Left, lastValue) {
		return false
	}

	if root != nil {
		if lastValue.init && lastValue.value >= root.Val {
			return false
		}
		lastValue.init = true
		lastValue.value = root.Val
	}

	return detour(root.Right, lastValue)
}
