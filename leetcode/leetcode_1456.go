package leetcode

func isVowel(vow string) bool {
	switch vow {
	case "a", "e", "i", "o", "u":
		return true
	default:
		return false
	}
}

func maxVowels(s string, k int) int {
	maxn, n, l := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if i < k {
			if isVowel(string(s[i])) {
				n++
				maxn++
			}

			continue
		}

		if n == k {
			return n
		}

		if isVowel(string(s[i])) {
			n++
		}

		if isVowel(string(s[l])) {
			n--
		}

		if n > maxn {
			maxn = n
			if maxn == k {
				return maxn
			}
		}

		l++
	}

	return maxn
}

func Leetcode1456(s string, k int) int {
	return maxVowels(s, k)
}
