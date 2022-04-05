package remove_vowels_from_a_string

// https://leetcode.com/problems/remove-vowels-from-a-string/

func removeVowels(s string) string {
	result := ""

	if len(s) == 0 {
		return result
	}

	exclude := []byte{'a', 'e', 'i', 'o', 'u'}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if contains(exclude, ch) {
			continue
		}

		result += string(ch)
	}

	return result
}

func contains(items []byte, item byte) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}

	return false
}
