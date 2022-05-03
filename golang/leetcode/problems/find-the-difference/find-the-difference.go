package find_the_difference

// https://leetcode.com/problems/find-the-difference/

func findTheDifference(s string, t string) byte {
	letters := [28]byte{}
	lastIndex := len(s)
	for i := 0; i < lastIndex; i++ {
		letters[t[i]-'a']++
		letters[s[i]-'a']--
	}
	letters[t[lastIndex]-'a']++

	for i, letter := range letters {
		if letter != 0 {
			return byte(i + 'a')
		}
	}

	return 0
}
