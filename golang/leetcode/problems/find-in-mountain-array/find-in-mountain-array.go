package find_in_mountain_array

// https://leetcode.com/problems/find-in-mountain-array/

type MountainArray struct {
	counter int
	arr     []int
}

func (a *MountainArray) get(index int) int {
	a.counter++
	if a.counter >= 100 {
		panic("Time Limit Exceeded")
	}
	return a.arr[index]
}
func (a *MountainArray) length() int {
	return len(a.arr)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()
	right := length
	pivot := right / 2
	var left = 0

	cache := make(map[int]int, length/2)
	get := func(i int) int {
		if v, ok := cache[i]; ok {
			return v
		}

		v := mountainArr.get(i)
		cache[i] = v
		return v
	}

	if get(left) == target {
		return left
	}

	for {
		if get(pivot-1) < get(pivot) && get(pivot) > get(pivot+1) {
			break
		}

		if get(pivot-1) < get(pivot) && get(pivot) < get(pivot+1) {
			left, pivot = pivot, pivot+(right-pivot)/2
			continue
		}

		if get(pivot-1) > get(pivot) && get(pivot) > get(pivot+1) {
			right, pivot = pivot, left+(pivot-left)/2
			continue
		}
	}

	var highestIndex = pivot + 1

	left = 0
	right = highestIndex
	pivot = right / 2
	for pivot > left {
		v := get(pivot)
		if v == target {
			return pivot
		} else if v > target {
			right = pivot
		} else {
			left = pivot
		}
		pivot = left + (right-left)/2
	}
	if get(pivot) == target {
		return pivot
	}

	left = highestIndex
	right = length
	pivot = left + (right-left)/2
	for pivot > left {
		v := get(pivot)
		if v == target {
			return pivot
		} else if v < target {
			right = pivot
		} else {
			left = pivot
		}
		pivot = left + (right-left)/2
	}
	if get(pivot) == target {
		return pivot
	}

	return -1
}
