package leetcode

func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	if len1 == 0 || len2 == 0 {
		return ""
	}

	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[0:gcd(len1, len2)]
}

func gcd(len1, len2 int) int {
	for len2 > 0 {
		len1, len2 = len2, len1%len2
	}

	return len1
}

func Leetcode1071(str1 string, str2 string) string {
	return gcdOfStrings(str1, str2)
}
