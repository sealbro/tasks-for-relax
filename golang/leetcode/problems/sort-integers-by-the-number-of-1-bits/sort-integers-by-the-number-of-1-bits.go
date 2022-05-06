package sort_integers_by_the_number_of_1_bits

// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/

func sortByBits(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	bits := make(map[int]int)

	return mergeSort(bits, arr)
}

func mergeSort(bits map[int]int, arr []int) []int {
	pivot := len(arr) / 2

	left := arr[:pivot]
	right := arr[pivot:]

	if len(left) > 1 {
		left = mergeSort(bits, left)
	}

	if len(right) > 1 {
		right = mergeSort(bits, right)
	}

	var sorted []int

	for i, j := 0, 0; i < len(left) || j < len(right); {
		if j >= len(right) || i < len(left) && compare(bits, left[i], right[j]) {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	return sorted
}

func compare(bits map[int]int, left, right int) bool {
	one := bitsCount(bits, left)
	two := bitsCount(bits, right)
	ok := false
	if one == two {
		ok = left < right
	} else {
		ok = one < two
	}
	return ok
}

func bitsCount(bits map[int]int, num int) int {
	if count, ok := bits[num]; ok {
		return count
	}

	count := 0
	for i := 0; i < 18; i++ {
		if num>>i&1 == 1 {
			count++
		}
	}
	bits[num] = count

	return count
}
