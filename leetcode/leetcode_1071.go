package leetcode

func gcdOfStrings(str1 string, str2 string) string {
	gcd := func(number1, number2 int) int {
		for number2 > 0 {
			number1, number2 = number2, number1%number2
		}

		return number1
	}

	len1 := len(str1)
	len2 := len(str2)

	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[0:gcd(len1, len2)]
}

func Leetcode1071(str1 string, str2 string) string {
	return gcdOfStrings(str1, str2)
}
