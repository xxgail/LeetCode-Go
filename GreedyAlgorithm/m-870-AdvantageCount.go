package GreedyAlgorithm

import (
	"sort"
)

type SortInt []int

func (s SortInt) Len() int {
	return len(s)
}

func (s SortInt) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func (s SortInt) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s *SortInt) Delete(i int) {
	tmp1 := (*s)[:i]
	tmp2 := (*s)[i+1:]
	*s = append(tmp1, tmp2...)
}

/**
 * @Time: 2020/3/11 18:45
 * @DESC: 870. 田忌赛马 优势洗牌
 * 给定两个大小相等的数组 A 和 B，A 相对于 B 的优势可以用满足 A[i] > B[i] 的索引 i 的数目来描述。
 * 返回 A 的任意排列，使其相对于 B 的优势最大化。
 * @params A B
 * @return []int
 */
func AdvantangeCount(A []int, B []int) []int {
	s := SortInt(A)
	sort.Sort(s)

	var result []int

	for i := 0; i < len(B); i++ {
		start, end, mid := 0, s.Len()-1, 0
		for start < end {
			mid = (start + end) / 2
			if s[mid] > B[i] {
				end = mid
			} else {
				start = mid + 1
			}
		}

		if s[start] <= B[i] {
			start = 0
		}
		result = append(result, s[start])

		(&s).Delete(start)
	}
	return result
}
