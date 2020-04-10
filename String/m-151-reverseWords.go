package String

import "strings"

/**
 * @Time: 2020-4-10 13:47:50
 * @DESC: 151. 翻转字符串里的单词 中等
 * 给定一个字符串，逐个翻转字符串中的每个单词。
 * @param s
 * @return string
 */
func ReverseWords(s string) string {
	var res string
	var arr = strings.Split(s, " ")
	for _, v := range arr {
		if len(v) > 0 {
			res = v + " " + res
		}
	}
	if len(res) > 0 {
		res = res[0 : len(res)-1]
	}
	return res
}
