package leetcode

import (
	"strings"
)

func reverseWords(s string) string {
	strs := strings.Fields(s)

	p1 := 0
	p2 := len(strs) - 1

	for p1 < p2 {
		strs[p1], strs[p2] = strs[p2], strs[p1]
		p1++
		p2--
	}

	return strings.Join(strs, " ")
}

func Leetcode151(s string) string {
	return reverseWords(s)
}
