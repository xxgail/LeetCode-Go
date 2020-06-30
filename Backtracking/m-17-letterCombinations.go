package Backtracking

func LetterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}
	letterFunc("", digits, &result)
	return result
}

func letterFunc(res string, digits string, result *[]string) {
	dict := make(map[string][]string)
	dict["2"] = []string{"a", "b", "c"}
	dict["3"] = []string{"d", "e", "f"}
	dict["4"] = []string{"g", "h", "i"}
	dict["5"] = []string{"j", "k", "l"}
	dict["6"] = []string{"m", "n", "o"}
	dict["7"] = []string{"p", "q", "r", "s"}
	dict["8"] = []string{"t", "u", "v"}
	dict["9"] = []string{"w", "x", "y", "z"}
	if digits == "" {
		*result = append(*result, res)
		return
	}

	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		letterFunc(res, digits, result)
		res = res[0 : len(res)-1]
	}
}
