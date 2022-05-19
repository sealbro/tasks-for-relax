package the_k_weakest_rows_in_a_matrix

import "sort"

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

func kWeakestRows(mat [][]int, k int) []int {
	counts := make([]int, len(mat))

	for i, row := range mat {
		counts[i] = countSolders(row)
	}

	hashMap := make(map[int][]int)

	for i, val := range counts {
		if arr, ok := hashMap[val]; ok {
			hashMap[val] = append(arr, i)
		} else {
			hashMap[val] = []int{i}
		}
	}

	sort.Ints(counts)

	result := make([]int, k)
	j := 0
	for _, val := range counts {
		for _, i := range hashMap[val] {
			result[j] = i
			j++
			if k == j {
				return result
			}
		}
		hashMap[val] = nil
	}

	return result
}

func countSolders(row []int) int {
	l := 0
	r := len(row) - 1

	for l <= r {
		p := l + (r-l)/2

		if row[p] > 0 {
			l = p + 1
		} else {
			r = p - 1
		}
	}

	return l
}
