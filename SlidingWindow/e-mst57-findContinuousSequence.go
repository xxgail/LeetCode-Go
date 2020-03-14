package SlidingWindow

/**
 * @Time: 2020/3/13 22:24
 * @DESC: 面试题57 - II. 和为s的连续正数序列 简单
 * 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
 * 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
 * @params int
 * @return [][]int
 */
func FindContinuousSequence(target int) [][]int {
	var res [][]int
	left := 1
	right := 2
	for left < right && left <= target/2 {
		sum := (left + right) * (right - left + 1) / 2
		if sum == target {
			res = append(res, rangeArr(left, right, 1))
			left++
			right++
		} else if sum < target {
			right++
		} else {
			left++
		}
	}
	return res
}

// 参考PHP的range()方法
func rangeArr(start, end, step int) []int {
	step = 1
	var res []int
	for i := start; i <= end; i += step {
		res = append(res, i)
	}
	return res
}

// 滑动窗口
