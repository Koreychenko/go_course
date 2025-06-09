package leetcode

func predictPartyVictory(senate string) string {
	radiantQueue := make([]int, 0, len(senate))
	direQueue := make([]int, 0, len(senate))

	var direIndex, radiantIndex int

	for i, char := range senate {
		if char == 'R' {
			radiantQueue = append(radiantQueue, i)
		} else {
			direQueue = append(direQueue, i)
		}
	}

	for len(radiantQueue) > 0 && len(direQueue) > 0 {
		direIndex = direQueue[0]
		radiantIndex = radiantQueue[0]

		direQueue = direQueue[1:]
		radiantQueue = radiantQueue[1:]

		if direIndex < radiantIndex {
			direQueue = append(direQueue, direIndex+len(senate))
		} else {
			radiantQueue = append(radiantQueue, radiantIndex+len(senate))
		}
	}

	if len(radiantQueue) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func Leetcode649(senate string) string {
	return predictPartyVictory(senate)
}
