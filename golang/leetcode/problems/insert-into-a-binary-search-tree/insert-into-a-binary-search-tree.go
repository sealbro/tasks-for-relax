package insert_into_a_binary_search_tree

// https://leetcode.com/problems/insert-into-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	detour(root, val)

	return root
}

func detour(root *TreeNode, val int) {
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val: val,
			}
		} else {
			detour(root.Left, val)
		}
	} else if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: val,
			}
		} else {
			detour(root.Right, val)
		}
	}
}
