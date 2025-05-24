package leetcode

func maxOperations(nums []int, k int) int {
	hash := make(map[int]int)

	for _, v := range nums {
		if v >= k {
			/* As nums[i] can't be 0 */
			continue
		}

		if _, ok := hash[v]; !ok {
			hash[v] = 1
		} else {
			hash[v]++
		}
	}

	counter := 0
	for hk, hv := range hash {
		if hk == k/2 {
			if k-hk == hk {
				counter += hv / 2
				continue
			}
		}

		if hv2, ok := hash[k-hk]; ok {
			counter += min(hv, hv2)
			hash[k-hk] = 0
		}
	}

	return counter
}

func Leetcode1679(nums []int, k int) int {
	return maxOperations(nums, k)
}
