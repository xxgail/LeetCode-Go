package DynamicProgramming

import (
	"strings"
)

/**
 * @Time: 2020/3/19 23:04
 * @DESC: 139.单词拆分
 * @params
 * @return
 */
func WordBreak(s string, wordDict []string) bool {
	dp := make(map[string]bool)
	dp[""] = true
	for i := 1; i <= len(s); i++ {
		str := s[:i]
		dp[str] = false
		for _, v := range wordDict {
			k := strings.LastIndex(str, v)
			if k == -1 {
				continue
			}
			a := str[:k] + str[(k+len(v)):]
			_, ok := dp[a]
			dp[str] = ok && dp[a]
			if dp[str] == true {
				break
			}
		}
	}
	return dp[s]
}

// strings.LastIndex(str,substr string)
// 查找substr在str中最后一次出现的位置，返回坐标，如果没有就返回-1
