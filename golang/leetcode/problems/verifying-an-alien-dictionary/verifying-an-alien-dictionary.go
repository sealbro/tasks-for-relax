package verifying_an_alien_dictionary

import "math"

// https://leetcode.com/problems/verifying-an-alien-dictionary/

func isAlienSorted(words []string, order string) bool {
	alphabet := [26]byte{}
	for i, char := range order {
		alphabet[char-'a'] = byte(i)
	}

	for i := 0; i < len(words)-1; i++ {
		if !compareWords(words[i], words[i+1], alphabet) {
			return false
		}
	}

	return true
}

func compareWords(word1, word2 string, alphabet [26]byte) bool {
	length := int(math.Min(float64(len(word1)), float64(len(word2))))

	for i := 0; i < length; i++ {
		if alphabet[word1[i]-'a'] > alphabet[word2[i]-'a'] {
			return false
		}

		if alphabet[word1[i]-'a'] < alphabet[word2[i]-'a'] {
			return true
		}
	}

	return len(word1) <= len(word2)
}
