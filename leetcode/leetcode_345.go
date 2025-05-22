package leetcode

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	str := []byte(s)

	p1 := 0
	p2 := len(s) - 1
	var flag1, flag2 bool

	for p1 < p2 {
		if _, ok := vowels[str[p1]]; !ok {
			p1 += 1
			flag1 = false
		} else {
			flag1 = true
		}

		if _, ok := vowels[str[p2]]; !ok {
			p2 -= 1
			flag2 = false
		} else {
			flag2 = true
		}

		if flag1 && flag2 {
			str[p1], str[p2] = str[p2], str[p1]
			p1++
			p2--
		}
	}

	return string(str)
}

func Leetcode345(s string) string {
	return reverseVowels(s)
}
