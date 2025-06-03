package leetcode

func absLg(num1, num2 int) bool {
	return max(num1, -1*num1) > max(num2, -1*num2)
}

func absEq(num1, num2 int) bool {
	return max(num1, -1*num1) == max(num2, -1*num2)
}

func check(stack []int, v int) []int {
	ln := len(stack)
	if ln == 0 || stack[ln-1] < 0 || v > 0 {
		stack = append(stack, v)

		return stack
	}

	if v < 0 {
		if absEq(stack[ln-1], v) {
			stack = stack[0 : ln-1]

			return stack
		}

		if absLg(v, stack[ln-1]) {
			stack = stack[0 : ln-1]
			stack = check(stack, v)
		}
	}

	return stack
}

func asteroidCollision(asteroids []int) []int {
	var stack []int

	for _, v := range asteroids {
		stack = check(stack, v)
	}

	return stack
}

func Leetcode735(asteroids []int) []int {
	return asteroidCollision(asteroids)
}
