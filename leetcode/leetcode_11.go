package leetcode

func maxArea(height []int) int {
	sqr, p1, p2 := 0, 0, len(height)-1

	for p1 != p2 {
		sqr = max(sqr, (p2-p1)*min(height[p1], height[p2]))

		if height[p1] < height[p2] {
			p1++

			continue
		}

		p2--
	}

	return sqr
}

func Leetcode11(height []int) int {
	return maxArea(height)
}
