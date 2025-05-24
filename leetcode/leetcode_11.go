package leetcode

func maxArea(height []int) int {
	p1 := 0
	p2 := len(height) - 1

	sqr := 0

	for p1 != p2 {
		sqrCandidate := (p2 - p1) * min(height[p1], height[p2])
		if sqr < sqrCandidate {
			sqr = sqrCandidate
		}

		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return sqr
}

func Leetcode11(height []int) int {
	return maxArea(height)
}
