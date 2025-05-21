package leetcode

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxValue := slices.Max(candies)
	result := make([]bool, len(candies))

	for i, val := range candies {
		result[i] = val+extraCandies >= maxValue
	}

	return result
}

func Leetcode1431(candies []int, extraCandies int) []bool {
	return kidsWithCandies(candies, extraCandies)
}
