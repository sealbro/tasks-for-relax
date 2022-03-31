package longest_substring_without_repeating_characters

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}

	start := 0
	maxCount := 1
	for i := 1; i < length; i++ {
		count := 1
		if s[i-1] == s[i] {
			start = i
			continue
		}

		for j := start; j < i; j++ {
			if s[i] == s[j] {
				start = j + 1
				break
			}
			count++
		}

		if maxCount < count {
			maxCount = count
		}
	}

	return maxCount
}
