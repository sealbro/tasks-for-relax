package can_place_flowers

// https://leetcode.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 1 {
		return flowerbed[0] == 0 // 0 <= n <= flowerbed.length
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		n--
	}

	if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		flowerbed[len(flowerbed)-1] = 1
		n--
	}

	for i := 1; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i-2] == 0 {
			n--
			flowerbed[i-1] = 1
		}
	}

	return n <= 0
}
