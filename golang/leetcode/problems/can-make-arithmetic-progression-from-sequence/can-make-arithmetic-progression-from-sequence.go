package can_make_arithmetic_progression_from_sequence

import "sort"

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if diff != arr[i+1]-arr[i] {
			return false
		}
	}

	return true
}
