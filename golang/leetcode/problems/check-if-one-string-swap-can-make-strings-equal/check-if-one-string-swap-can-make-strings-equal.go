package check_if_one_string_swap_can_make_strings_equal

// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/

func areAlmostEqual(s1 string, s2 string) bool {
	letters := [28]byte{}
	diffCount := 0
	for i := 0; i < len(s1); i++ {
		letters[s1[i]-'a']++
		letters[s2[i]-'a']--
		if s1[i] != s2[i] {
			diffCount++
		}

		if diffCount > 2 {
			return false
		}
	}

	for _, count := range letters {
		if count != 0 {
			return false
		}
	}

	return true
}
