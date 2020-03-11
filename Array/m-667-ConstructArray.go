package Array

/**
 * @Time: 2019/9/3 18:49
 * @DESC: 667. 优美的排列② 中等
 * 给定两个整数 n 和 k，你需要实现一个数组，这个数组包含从 1 到 n 的 n 个不同整数，同时满足以下条件：
 * ① 如果这个数组是 [a1, a2, a3, ... , an] ，那么数组 [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] 中应该有且仅有 k 个不同整数；.
 * ② 如果存在多种答案，你只需实现并返回其中任意一种.
 * @params n k
 * @return []int
 */
func ConstructArray(n int, k int) []int {
	data := []int{1}
	if k == 1 {
		for i := 1; i <= n; i++ {
			data = append(data, i)
		}
	} else {
		for i := 1; i <= k; i++ {
			if i%2 == 0 {
				data = append(data, (i+2)/2)
			} else {
				data = append(data, 1+k-(i-2)/2)
			}
		}

		for i := k + 1; i < n; i++ {
			data = append(data, i+1)
		}
	}
	return data
}
