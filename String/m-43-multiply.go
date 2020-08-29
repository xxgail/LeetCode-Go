package String

import "strconv"

/**
 * @Date 2020/8/29
 * @Desc 43. 字符串相乘
给定两个以字符串形式表示的非负整数num1和num2，返回num1和num2的乘积，它们的乘积也表示为字符串形式。
示例 1:
	输入: num1 = "2", num2 = "3"
	输出: "6"
示例2:
	输入: num1 = "123", num2 = "456"
	输出: "56088"
说明：
	num1 和 num2 的长度小于110。
	num1 和 num2 只包含数字0-9。
	num1 和 num2 均不以零开头，除非是数字 0 本身。
	不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 * @Param
 * @return
 * @link https://leetcode-cn.com/problems/multiply-strings
 **/
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1 := len(num1) - 1
	n2 := len(num2) - 1
	list := make([]int, n1+n2+2)
	n := n1 + n2 + 1
	for n1 >= 0 {
		j := n2
		for j >= 0 {
			n := int(num1[n1]-'0') * int(num2[j]-'0')
			list[n1+j+1] += n
			j--
		}
		n1--
	}
	result := ""
	add := 0
	for n >= 0 || add != 0 {
		s := list[n] + add
		add = s / 10
		result = strconv.Itoa(s%10) + result
		n--
	}
	if result[0] == '0' {
		return result[1:]
	} else {
		return result
	}
}
