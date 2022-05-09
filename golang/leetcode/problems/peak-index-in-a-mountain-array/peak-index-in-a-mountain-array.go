package peak_index_in_a_mountain_array

// https://leetcode.com/problems/peak-index-in-a-mountain-array/

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr)
	pivot := end / 2

	for {
		if arr[pivot-1] < arr[pivot] && arr[pivot] > arr[pivot+1] {
			break
		}

		if arr[pivot-1] < arr[pivot] && arr[pivot] < arr[pivot+1] {
			start, pivot = pivot, pivot+(end-pivot+1)/2
			continue
		}

		if arr[pivot-1] > arr[pivot] && arr[pivot] > arr[pivot+1] {
			end, pivot = pivot, start+(pivot-start+1)/2
			continue
		}
	}

	return pivot
}
