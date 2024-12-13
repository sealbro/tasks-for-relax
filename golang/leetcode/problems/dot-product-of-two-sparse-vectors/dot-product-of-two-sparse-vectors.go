package dot_product_of_two_sparse_vectors

// https://leetcode.com/problems/dot-product-of-two-sparse-vectors/

type SparseVector struct {
	nums []int
}

func Constructor(nums []int) SparseVector {
	return SparseVector{nums: nums}
}

// Return the dotProduct of two sparse vectors
func (v *SparseVector) dotProduct(vec SparseVector) int {
	result := 0
	for i, num := range v.nums {
		result += num * vec.nums[i]
	}

	return result
}
