package Stack

/**
 * @Date 2020/8/29
 * @Desc 20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串，判断字符串是否有效。
有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	注意空字符串可被认为是有效字符串。
示例 1:
	输入: "()"
	输出: true
 * @Param s string
 * @return bool
 * @link https://leetcode-cn.com/problems/valid-parentheses
 **/
func IsValid(s string) bool {
	lenS := len(s)
	var stack []byte
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < lenS; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == m[s[i]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
