package kids_with_the_greatest_number_of_candies

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	results := make([]bool, len(candies))
	maxC := 0
	for _, v := range candies {
		if maxC < v {
			maxC = v
		}
	}

	for i, v := range candies {
		results[i] = v+extraCandies >= maxC
	}

	return results
}
