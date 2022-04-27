package two_sum_iv_input_is_a_bst

// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var values []int
	detour(root, &values)

	current := 0

	i := 0
	for current < len(values) {
		if i != current && values[i] == k-values[current] {
			return true
		}

		i++

		if i == len(values) {
			current++
			i = 0
		}
	}

	return false
}

func detour(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}

	detour(root.Left, values)
	*values = append(*values, root.Val)
	detour(root.Right, values)
}
