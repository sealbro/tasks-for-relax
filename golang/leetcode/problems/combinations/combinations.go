package combinations

import "fmt"

// https://leetcode.com/problems/combinations/

func combine(n int, k int) [][]int {
	var output [][]int

	var nums []int
	for i := 1; i < k+1; i++ {
		nums = append(nums, i)
	}
	nums = append(nums, n+1)

	j := 0

	for j < k {

		fmt.Printf("%v %v\n", nums, j)
		output = append(output, append([]int{}, nums[:k]...))
		j = 0
		for j < k && nums[j+1] == nums[j]+1 {
			nums[j] = j + 1
			j += 1
			fmt.Printf("\t%v %v\n", nums, j)
		}
		nums[j] += 1
	}

	return output
}
