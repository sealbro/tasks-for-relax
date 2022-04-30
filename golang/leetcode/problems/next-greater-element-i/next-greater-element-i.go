package next_greater_element_i

// https://leetcode.com/problems/next-greater-element-i/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int
	for i := 0; i < len(nums1); i++ {
		num := -1

		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				for g := j + 1; g < len(nums2); g++ {
					if nums2[j] < nums2[g] {
						num = nums2[g]
						break
					}
				}

				break
			}
		}

		result = append(result, num)
	}

	return result
}
