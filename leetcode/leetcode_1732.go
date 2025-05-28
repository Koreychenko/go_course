package leetcode

func largestAltitude(gain []int) int {
	height, maxHeight := 0, 0

	for i := 0; i < len(gain); i++ {
		height += gain[i]
		maxHeight = max(maxHeight, height)
	}

	return maxHeight
}

func Leetcode1732(gain []int) int {
	return largestAltitude(gain)
}
