package concatenation_of_array

// https://leetcode.com/problems/concatenation-of-array/

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i := range nums {
		ans[i+n] = nums[i]
		ans[i] = nums[i]
	}

	return ans
}
