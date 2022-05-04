package decrypt_string_from_alphabet_to_integer_mapping

import "math"

// https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/

// 97 - 105
//Characters ('a' to 'i') are represented by ('1' to '9') respectively.
// 106 - 122
//Characters ('j' to 'z') are represented by ('10#' to '26#') respectively.

func freqAlphabets(s string) string {
	output := ""
	isSharp := false
	const empty = math.MaxInt8
	var buffer byte = empty
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			isSharp = true
			continue
		}

		if isSharp {
			if buffer == empty {
				buffer = s[i] - '0'
			} else {
				buffer += (s[i]-'0')*10 - 10
				output = string('j'+buffer) + output
				buffer = empty
				isSharp = false
			}
			continue
		}

		output = string('a'+(s[i]-'1')) + output
	}

	return output
}
