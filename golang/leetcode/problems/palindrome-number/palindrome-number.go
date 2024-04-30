package palindrome_number

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := 0
	step := x
	for {
		num = num + (step % 10)
		step = step / 10
		if step != 0 {
			num = num * 10
		} else {
			break
		}
	}

	return num == x
}
