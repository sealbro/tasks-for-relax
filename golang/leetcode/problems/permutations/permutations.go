package permutations

import (
	"sort"
)

// https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	var permuted [][]int

	backtrack(&permuted, nums, 0)

	return permuted
}

func backtrack(permuted *[][]int, nums []int, first int) {
	nums = append([]int{}, nums...)

	if first >= len(nums)-1 {
		*permuted = append(*permuted, nums)
		return
	}

	for i := first; i < len(nums); i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backtrack(permuted, nums, first+1)
	}
}

func permuteLexicographic(nums []int) [][]int {
	sort.Ints(nums)
	permuted := [][]int{append([]int{}, nums...)}

	if len(nums) == 1 {
		return permuted
	}

	for {
		x := -1

		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				x = i
			}
		}

		if x < 0 {
			break
		}

		y := -1
		for i := 0; i < len(nums); i++ {
			if nums[x] < nums[i] {
				y = i
			}
		}

		if x < y {
			nums[x], nums[y] = nums[y], nums[x]

			firstPart := append([]int{}, nums[:x+1]...)
			secondPart := append([]int{}, nums[x+1:]...)

			reverseIndex := len(secondPart) - 1
			for i := 0; i < len(secondPart)/2; i++ {
				secondPart[i], secondPart[reverseIndex] = secondPart[reverseIndex], secondPart[i]
				reverseIndex--
			}

			nums = append(firstPart, secondPart...)
		}

		permuted = append(permuted, append([]int{}, nums...))
	}

	return permuted
}
