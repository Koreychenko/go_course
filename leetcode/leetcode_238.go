package leetcode

func productExceptSelf(nums []int) []int {
	ln := len(nums)

	var left []int
	var right []int

	k := 0
	for k < ln {
		left = append(left, 1)
		right = append(right, 1)
		k++
	}

	for i := 1; i < ln; i++ {
		left[i] = nums[i-1] * left[i-1]
		right[ln-i-1] = nums[ln-i] * right[ln-i]
	}

	var result []int

	for i := 0; i < ln; i++ {
		result = append(result, left[i]*right[i])
	}

	return result
}

func Leetcode238(nums []int) []int {
	return productExceptSelf(nums)
}
