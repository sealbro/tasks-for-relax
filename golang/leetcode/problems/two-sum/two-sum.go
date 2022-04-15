package two_sum

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	mapkins := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		diff := target - num

		if index, ok := mapkins[diff]; ok {
			return []int{index, i}
		} else {
			mapkins[num] = i
		}
	}

	return []int{}
}
