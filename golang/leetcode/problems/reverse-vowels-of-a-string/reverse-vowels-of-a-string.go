package reverse_vowels_of_a_string

// https://leetcode.com/problems/reverse-vowels-of-a-string/

func reverseVowels(s string) string {
	if len(s) == 1 {
		return s
	}

	vowels := map[byte]struct{}{
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	output := []byte(s)

	si := 0
	li := len(s) - 1
	for si < li {
		for si < len(s) {
			_, ok := vowels[output[si]]
			if ok {
				break
			}

			si++
		}
		for li >= 0 {
			_, ok := vowels[output[li]]
			if ok {
				break
			}

			li--
		}

		if si < len(s) && li >= 0 && si < li {
			output[si], output[li] = output[li], output[si]
			si++
			li--
		}
	}

	return string(output)
}
