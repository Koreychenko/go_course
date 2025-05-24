package leetcode

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

func Leetcode283(nums []int) {
	moveZeroes(nums)
}
