package find_resultant_array_after_removing_anagrams

// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/

func removeAnagrams(words []string) []string {
	if len(words) < 2 {
		return words
	}

	l, r := 0, 1

	for r < len(words) {
		if len(words[l]) != len(words[r]) {
			l = r
			r++
			continue
		}

		diff := [26]int{}

		for i := 0; i < len(words[l]); i++ {
			diff[words[l][i]-'a']++
			diff[words[r][i]-'a']--
		}

		found := true
		for _, c := range diff {
			if c != 0 {
				l = r
				r++
				found = false
				break
			}
		}

		if found {
			words[r] = ""
			r++
		}
	}

	var result []string

	for _, v := range words {
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}
