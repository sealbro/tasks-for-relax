package find_the_length_of_the_longest_common_prefix

// https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	arr1map := make(map[int]bool)
	for _, n1 := range arr1 {
		if _, ok := arr1map[n1]; ok {
			continue
		}
		for n1 != 0 {
			arr1map[n1] = true
			n1 /= 10
		}
	}

	numLen := func(n int) int {
		c := 0

		for n != 0 {
			c++
			n /= 10
		}

		return c
	}

	maxSeq := 0
	for _, n2 := range arr2 {
		for n2 != 0 {
			if _, ok := arr1map[n2]; ok {
				newLen := numLen(n2)
				if newLen > maxSeq {
					maxSeq = newLen
				}
				break
			}
			n2 /= 10
		}
	}

	return maxSeq
}
