package String

/**
680. 验证回文字符串 Ⅱ
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
https://leetcode-cn.com/problems/valid-palindrome-ii/
*/
func ValidPalindrome(s string) bool {
	if validPalindromeFunc(s) == true {
		return true
	}
	l := 0
	r := len(s) - 1
	for true {
		if s[l] == s[r] {
			l++
			r--
		} else {
			break
		}
	}
	return validPalindromeFunc(s[:l]+s[l+1:]) || validPalindromeFunc(s[:r]+s[r+1:])
}

func validPalindromeFunc(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
