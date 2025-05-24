package leetcode

func moveZeroesFast(nums []int) []int {
	prev := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != prev {
				nums[prev] = nums[i]
				nums[i] = 0
			}

			prev++
		}
	}

	return nums
}

func Leetcode283(nums []int) []int {
	return moveZeroesFast(nums)
}
