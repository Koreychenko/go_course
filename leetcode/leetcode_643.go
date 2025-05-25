package leetcode

import "math"

func findMaxAverage(nums []int, k int) float64 {
	sumK := 0

	avg := math.Inf(-1)

	for i := 0; i < len(nums); i++ {
		if i < k {
			sumK += nums[i]
		}

		if i == k {
			avg = max(avg, float64(sumK)/float64(k))
		}

		if i >= k {
			sumK = sumK - nums[i-k] + nums[i]
			avg = max(avg, float64(sumK)/float64(k))
		}
	}

	avg = max(avg, float64(sumK)/float64(k))

	return avg
}

func Leetcode643(nums []int, k int) float64 {
	return findMaxAverage(nums, k)
}
