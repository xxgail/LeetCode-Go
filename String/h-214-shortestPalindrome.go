package String

/**
 * @Date 2020/8/29
 * @Desc 214. 最短回文串
 * 给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
 * 示例1:
	输入: "aacecaaa"
	输出: "aaacecaaa"
 * 示例 2:
	输入: "abcd"
	输出: "dcbabcd"
 * @Param s string
 * @return s string
 * @link https://leetcode-cn.com/problems/shortest-palindrome
 **/
func ShortestPalindrome(s string) string {
	l := len(s) - 1
	i := 0
	for !isPalindrome(s) {
		s = s[:i] + s[l:l+1] + s[i:]
		i++
	}
	return s
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
