package average_salary_excluding_the_minimum_and_maximum_salary

import "math"

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

func average(salary []int) float64 {
	min := math.MaxInt
	max := 0
	sum := 0
	count := float64(len(salary) - 2)

	for _, sal := range salary {
		if sal > max {
			max = sal
		}

		if sal < min {
			min = sal
		}

		sum += sal
	}

	return float64(sum-max-min) / count
}
