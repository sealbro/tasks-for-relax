package two_sum_ii_input_array_is_sorted

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum(numbers []int, target int) []int {
	result := []int{0, 0}
	for i := 0; i < len(numbers); i++ {
		result[0] = i + 1
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				result[1] = j + 1

				return result
			}
		}
	}
	return result
}
