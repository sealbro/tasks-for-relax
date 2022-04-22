package valid_parentheses

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	var stack []int32

	lastIndex := -1
	for _, ch := range s {
		if lastIndex >= 0 && ch-stack[lastIndex] > 0 && ch-stack[lastIndex] < 3 {
			lastIndex--
		} else {
			lastIndex++
			if len(stack) <= lastIndex {
				stack = append(stack, ch)
			} else {
				stack[lastIndex] = ch
			}

		}
	}

	return lastIndex < 0
}
