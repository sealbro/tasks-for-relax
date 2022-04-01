package valid_anagram

// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	sLength := len(s)
	tLength := len(t)

	if sLength != tLength {
		return false
	}

	charMap := [26]int{}
	for i := 0; i < sLength; i++ {
		charMap[(s[i]-'a')]++
		charMap[(t[i]-'a')]--
	}

	for _, m := range charMap {
		if m != 0 {
			return false
		}
	}

	return true
}
