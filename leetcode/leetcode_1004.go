package leetcode

func longestOnes(nums []int, k int) int {
	maxln := 0 // Max length
	n := 0     // Number of found 0
	l := 0     // Left pointer
	z := 0
	for i := 0; i < len(nums); i++ {
		z = i
		if nums[i] == 0 {
			n++
		}
		if n == k {
			maxln = max(maxln, i-l+1)
		}

		if n > k {
			for {
				if nums[l] == 0 {
					n--
					l++
					break
				}
				l++
			}
		}
	}

	return max(maxln, z-l+1)
}

func Leetcode1004(nums []int, k int) int {
	return longestOnes(nums, k)
}
