package String

/**
 * @Time: 2020/3/12 17:53
 * @DESC: 1071. 字符串的最大公因子 简单
 * 对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。
 * 返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。
 * @param str1
 * @param str2
 * @return string
 */
func GocOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[0:goc(len(str1), len(str2))] // go的字符串截取可以类似于数组直接截取
}

func goc(a, b int) int {
	if b == 0 {
		return a
	} else {
		return goc(b, a%b)
	}
}
