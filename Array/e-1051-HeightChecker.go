package Array

import "sort"

/**
 * @Time: 2019/9/3 18:53
 * @DESC: 1051. 高度检查器 简单
 * 学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。
 * 请你返回至少有多少个学生没有站在正确位置数量。
 * 该人数指的是：能让所有学生以 非递减 高度排列的必要移动人数。
 * @params heights
 * @return int
 */
func HeightChecker(heights []int) int {
	NewHeight := make([]int, 0, len(heights)/2)

	a := 0
	for i := 0; i < len(heights); i++ {
		NewHeight = append(NewHeight, heights[i])
	}
	sort.Ints(NewHeight)
	for i := 0; i < len(heights); i++ {
		if heights[i]^NewHeight[i] != 0 {
			a = a + 1
		}
	}
	return a
}
