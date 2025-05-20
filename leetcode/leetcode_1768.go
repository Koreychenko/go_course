package leetcode

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var merged strings.Builder

	len1 := len(word1)
	len2 := len(word2)

	if len1 == 0 {
		return word2
	}

	if len2 == 0 {
		return word1
	}

	for i := 0; i < len1; i++ {
		merged.WriteByte(word1[i])
		if i < len2 {
			merged.WriteByte(word2[i])
		} else {
			merged.WriteString(word1[i+1:])
			break
		}
	}

	if len1 < len2 {
		merged.WriteString(word2[len1:])
	}

	return merged.String()
}

func Leetcode1768(word1 string, word2 string) string {
	return mergeAlternately(word1, word2)
}
