package first_unique_character_in_a_string

// https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	chars := [28]int{}

	index := 0
	for i, ch := range s {
		chars[ch-'a']++

		for chars[s[index]-'a'] > 1 && index < i {
			index++
		}
	}

	if chars[s[index]-'a'] > 1 {
		return -1
	}

	return index
}
