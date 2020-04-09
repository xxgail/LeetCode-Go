package Backtracking

/**
 * @Time: 2020/4/9 22:04
 * @DESC: 22. 括号生成
 * @params
 * @return
 */
func GenerateParenthesis(n int) []string {
	var result []string
	generateParenthesisFunc("", n, n, &result)
	return result
}

func generateParenthesisFunc(str string, left, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, str)
		return
	}

	if left > 0 {
		generateParenthesisFunc(str+"(", left-1, right, result)
	}

	if right > left {
		generateParenthesisFunc(str+")", left, right-1, result)
	}
}
