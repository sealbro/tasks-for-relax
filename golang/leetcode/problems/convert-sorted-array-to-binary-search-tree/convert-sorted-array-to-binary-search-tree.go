package convert_sorted_array_to_binary_search_tree

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return Fill(nums, 0, len(nums)-1)
}

func Fill(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	pivot := (start + end) / 2

	node := &TreeNode{
		Val: nums[pivot],
	}

	node.Left = Fill(nums, start, pivot-1)
	node.Right = Fill(nums, pivot+1, end)

	return node
}
