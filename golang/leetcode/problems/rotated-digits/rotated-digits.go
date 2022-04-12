package rotated_digits

// https://leetcode.com/problems/rotated-digits/

func rotatedDigits(n int) int {
	count := 0

	for i := 1; i < n+1; i++ {
		num := i
		isGood := true

		mult := 1
		newNum := 0
		for num > 0 {
			digit := num % 10

			if digit == 3 || digit == 4 || digit == 7 {
				isGood = false
				break
			}

			switch digit {
			case 2:
				newNum += 5 * mult
			case 5:
				newNum += 2 * mult
			case 6:
				newNum += 9 * mult
			case 9:
				newNum += 6 * mult
			default:
				newNum += digit * mult
			}

			num = num / 10
			mult *= 10
		}
		if isGood && i != newNum {
			count++
		}
	}

	return count
}
