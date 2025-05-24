package leetcode

func isSubsequence(s string, t string) bool {
	var sp int // There is a pointer for the substring

	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	for i := 0; i < len(t); i++ { // Iterate over the haystack
		if s[sp] == t[i] {
			sp++

			if sp == len(s) {
				return true
			}
		}
	}

	return false
}

func Leetcode392(s string, t string) bool {
	return isSubsequence(s, t)
}
