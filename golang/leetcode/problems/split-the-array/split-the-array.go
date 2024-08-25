package split_the_array

// https://leetcode.com/problems/split-the-array/

func isPossibleToSplit(nums []int) bool {
	numMap := make(map[int]int)
	for _, num := range nums {
		if v, ok := numMap[num]; ok {
			if v == 2 {
				return false
			}
			numMap[num] = v + 1
		} else {
			numMap[num] = 1
		}
	}

	return true
}
