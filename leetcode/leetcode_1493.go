package leetcode

func longestSubarray(nums []int) int {
	maxln := 0 // Max length
	n := 0     // Number of found 0
	l := 0     // Left pointer
	z := 0
	for i := 0; i < len(nums); i++ {
		z = i
		if nums[i] == 0 {
			n++
		}
		if n == 1 {
			maxln = max(maxln, i-l)
		}

		if n > 1 {
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

	return max(maxln, z-l)
}

func Leetcode1493(nums []int) int {
	return longestSubarray(nums)
}
