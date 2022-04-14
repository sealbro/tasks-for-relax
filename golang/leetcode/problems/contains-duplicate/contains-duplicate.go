package contains_duplicate

// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = struct{}{}
		} else {
			return true
		}
	}

	return false
}
