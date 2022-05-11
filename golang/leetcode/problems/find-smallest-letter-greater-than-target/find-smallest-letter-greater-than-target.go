package find_smallest_letter_greater_than_target

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/

func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[0] > target || letters[len(letters)-1] <= target {
		return letters[0]
	}

	l := 0
	r := len(letters) - 1

	for l <= r {
		p := l + (r-l)/2

		if p > 0 && letters[p-1] == target && letters[p] != target {
			return letters[p]
		}

		if letters[p] <= target {
			l = p + 1
		} else {
			r = p - 1
		}
	}

	return letters[l]
}
