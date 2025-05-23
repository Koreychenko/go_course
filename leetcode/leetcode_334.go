package leetcode

import "math"

func increasingTriplet(nums []int) bool {
	n1 := math.MaxInt32
	n2 := math.MaxInt32

	for _, n := range nums {
		if n <= n1 {
			n1 = n
		} else if n <= n2 {
			n2 = n
		} else {
			return true
		}
	}

	return false
}

func Leetcode334(nums []int) bool {
	return increasingTriplet(nums)
}
