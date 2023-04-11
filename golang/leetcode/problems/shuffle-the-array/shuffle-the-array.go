package shuffle_the_array

// https://leetcode.com/problems/shuffle-the-array/

func shuffle(nums []int, n int) []int {
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			nums[i+j], nums[j+n] = nums[j+n], nums[i+j]
		}
	}

	return nums
}
