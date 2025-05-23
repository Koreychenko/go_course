package leetcode

import (
	"slices"
	"strconv"
)

func compress(chars []byte) int {
	counter := 0
	i := 0

	for {
		if i != 0 {
			if chars[i] != chars[i-1] {
				if counter > 1 {
					counterBytes := []byte(strconv.Itoa(counter))
					chars = slices.Replace(chars, i-counter+1, i, counterBytes...)
					i = i - (counter - 1 - len(counterBytes))
				}

				counter = 0
			}
		}

		i++
		counter++
		if i == len(chars) {
			break
		}
	}

	if counter > 1 {
		counterBytes := []byte(strconv.Itoa(counter))
		chars = chars[0 : len(chars)-counter+1]
		chars = append(chars, counterBytes...)
	}

	return len(chars)
}

func Leetcode443(chars []byte) int {
	return compress(chars)
}
