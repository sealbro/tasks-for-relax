package sum_of_all_odd_length_subarrays

// https://leetcode.com/problems/sum-of-all-odd-length-subarrays/

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	for i := 1; i <= len(arr); i += 2 {
		for j := 0; j+i <= len(arr); j++ {
			for _, val := range arr[j : j+i] {
				sum += val
			}
		}
	}
	return sum
}
