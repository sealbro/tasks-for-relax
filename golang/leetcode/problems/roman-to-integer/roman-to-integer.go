package roman_to_integer

// https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
	symbol := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := symbol[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		current := symbol[s[i]]
		if current < symbol[s[i+1]] {
			result -= current
		} else {
			result += current
		}
	}

	return result
}
