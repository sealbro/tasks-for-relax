package range_sum_query_immutable

// https://leetcode.com/problems/range-sum-query-immutable/

type NumArray struct {
	nums  []int
	cache []int
}

func Constructor(nums []int) NumArray {
	cache := make([]int, len(nums)+1)

	//cache[0] = nums[0]

	for i := 0; i < len(nums); i++ {
		cache[i+1] = cache[i] + nums[i]
	}

	return NumArray{
		cache: cache,
		nums:  nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.cache[right+1] - this.cache[left]
}
