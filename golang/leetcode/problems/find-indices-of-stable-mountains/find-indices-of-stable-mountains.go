package find_indices_of_stable_mountains

// https://leetcode.com/problems/find-indices-of-stable-mountains/

func stableMountains(height []int, threshold int) []int {
	var indexes []int
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			indexes = append(indexes, i)
		}
	}

	return indexes
}
