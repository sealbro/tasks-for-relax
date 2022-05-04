package to_lower_case

// https://leetcode.com/problems/to-lower-case/

func toLowerCase(s string) string {
	var lowerStr = []byte(s)

	for i, char := range lowerStr {
		if char > 64 && char < 91 {
			lowerStr[i] = char + 32
		}
	}

	return string(lowerStr)
}
