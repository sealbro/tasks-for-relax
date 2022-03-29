package reverse_words_in_a_string_iii

// https://leetcode.com/problems/reverse-words-in-a-string-iii/

func reverseWords(s string) string {
	length := len(s)

	if length == 1 {
		return s
	}

	reversed := ""
	currentWord := ""

	for _, ch := range s {
		if ch == ' ' {
			reversed += currentWord + " "

			currentWord = ""
			continue
		}

		currentWord = string(ch) + currentWord
	}

	reversed += currentWord

	return reversed
}
