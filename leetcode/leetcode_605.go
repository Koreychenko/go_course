package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	var counter int

	if n == 0 {
		return true
	}

	flowerbedLen := len(flowerbed)

	if flowerbedLen == 1 && flowerbed[0] == 0 {
		return true
	}

	for i, v := range flowerbed {
		if i == 0 && v == 0 {
			counter++
		}

		if v == 0 {
			counter++
			if i == flowerbedLen-1 {
				counter++
			}
		} else {
			n -= (counter - 1) / 2
			counter = 0

			if n <= 0 {
				return true
			}
		}
	}

	if counter != 0 {
		n -= (counter - 1) / 2
		counter = 0
	}

	if n <= 0 {
		return true
	}

	return false
}

func Leetcode605(flowerbed []int, n int) bool {
	return canPlaceFlowers(flowerbed, n)
}
