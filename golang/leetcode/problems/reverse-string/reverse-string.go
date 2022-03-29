package reverse_string

// https://leetcode.com/problems/reverse-string/

func reverseString(s []byte) {
	length := len(s)

	if length == 1 {
		return
	}

	j := length - 1
	for i := 0; i < length/2; i++ {
		if s[i] != s[j] {
			s[i], s[j] = s[j], s[i]
		}
		j--
	}
}
