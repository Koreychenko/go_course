package leetcode

import "slices"

func moveZeroes(nums []int) {
	for i, n := range nums {
		if n == 0 {
			counter := 0

			for j := i + 1; j < len(nums); j++ {
				if nums[j] == 0 {
					counter++
				} else {
					break
				}
			}

			if i+counter < len(nums)-1 {
				nums[i], nums[i+counter+1] = nums[i+counter+1], nums[i]
			}
		}
	}
}

func moveZeroesFast(nums []int) []int {
	ln := len(nums)
	for i, n := range nums {
		if n == 0 {
			counter := 0

			j := i
			for {
				if j < ln && nums[j] == 0 {
					counter++
					j++
				} else {
					break
				}
			}

			if i+counter < ln {
				nums = slices.Replace(nums, i, i+counter)
				j = 0
				for j < counter {
					nums = append(nums, 0)
					j++
				}
			} else {
				break
			}
		}
	}

	return nums
}

func Leetcode283(nums []int) []int {
	return moveZeroesFast(nums)
}
