package sign_of_the_product_of_an_array

// https://leetcode.com/problems/sign-of-the-product-of-an-array/

func arraySign(nums []int) int {
	negative := 0
	positive := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}

		if num > 0 {
			positive++
		} else {
			negative++
		}
	}

	if negative&1 == 1 {
		return -1
	} else {
		return 1
	}
}
