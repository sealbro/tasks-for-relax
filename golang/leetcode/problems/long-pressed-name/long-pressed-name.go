package long_pressed_name

// https://leetcode.com/problems/long-pressed-name/

func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}

	j := 0

	for i := 0; i < len(typed); i++ {
		if j < len(name) && name[j] == typed[i] {
			j++
		} else {
			if j == 0 || typed[i] != name[j-1] {
				return false
			}
		}
	}

	return j == len(name)
}
