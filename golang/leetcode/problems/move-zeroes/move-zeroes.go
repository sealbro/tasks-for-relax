package move_zeroes

// https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	length := len(nums)
	for i := length - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i; j < length && j+1 < length && nums[j+1] != 0; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
