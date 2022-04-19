package ransom_note

// https://leetcode.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	chars := [28]int{}

	for _, ch := range magazine {
		chars[ch-'a']++
	}

	for _, ch := range ransomNote {
		chars[ch-'a']--

		if chars[ch-'a'] < 0 {
			return false
		}
	}

	return true
}
