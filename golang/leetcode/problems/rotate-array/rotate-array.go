package rotate_array

// https://leetcode.com/problems/rotate-array/

func rotate(nums []int, k int) {
	length := len(nums)

	if length == 1 || k == 0 || k == length {
		return
	}

	k = k % length

	count := 0
	for start := 0; count < length; start++ {
		current := start
		prev := nums[start]

		for {
			next := (current + k) % length
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++

			if start == current {
				break
			}
		}
	}
}
