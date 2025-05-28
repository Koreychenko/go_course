package leetcode

func pivotIndex(nums []int) int {
	ln := len(nums)
	leftSum := make([]int, ln)
	rightSum := make([]int, ln)

	for i := 0; i < ln; i++ {
		if i == 0 {
			leftSum[i] = 0
			rightSum[ln-1] = 0

			continue
		}

		leftSum[i] = leftSum[i-1] + nums[i-1]
		rightSum[ln-i-1] = rightSum[ln-i] + nums[ln-i]

		/* Optimization. In case left and right part are already calculated we can compare them */
		if i == ln-i-1 && leftSum[i] == rightSum[ln-i-1] {
			return i
		}
	}

	for i := 0; i < ln; i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}

	return -1
}

func Leetcode724(nums []int) int {
	return pivotIndex(nums)
}
