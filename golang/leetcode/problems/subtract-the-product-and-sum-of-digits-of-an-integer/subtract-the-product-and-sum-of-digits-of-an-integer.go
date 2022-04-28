package subtract_the_product_and_sum_of_digits_of_an_integer

// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/

func subtractProductAndSum(n int) int {
	sum := 0
	mult := 1

	for n > 0 {
		digit := n % 10
		sum += digit
		mult *= digit

		n = n / 10
	}

	return mult - sum
}
