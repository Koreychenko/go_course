package leetcode

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	newline := ""
	counter := 0
	leftPointer := 0
	multiplier := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' && counter == 0 {
			newline += string(s[i])
		} else if s[i] == '[' {
			if counter == 0 {
				leftPointer = i
			}
			counter++
		} else if s[i] == ']' {
			if counter != 1 {
				counter--
			} else {
				counter = 0
				mult, _ := strconv.Atoi(multiplier)
				newline += strings.Repeat(decodeString(s[leftPointer+1:i]), mult)
				multiplier = ""
			}
		} else {
			if counter == 0 {
				multiplier += string(s[i])
			}
		}
	}

	return newline
}

func Leetcode394(s string) string {
	return decodeString(s)
}
