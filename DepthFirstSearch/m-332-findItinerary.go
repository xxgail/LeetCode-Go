package DepthFirstSearch

import "sort"

var path []string
var images map[string][]string

/**
 * @Date 2020/8/29
 * @Desc 332. 重新安排行程
给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。
提示：
	如果存在多种有效的行程，请你按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
	所有的机场都用三个大写字母表示（机场代码）。
	假定所有机票至少存在一种合理的行程。
	所有的机票必须都用一次 且 只能用一次。
示例 1：
	输入：[["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
	输出：["JFK", "MUC", "LHR", "SFO", "SJC"]
 * @Param tickets [][]string
 * @return []string
 * @link https://leetcode-cn.com/problems/reconstruct-itinerary/
 **/
func FindItinerary(tickets [][]string) []string {
	images = map[string][]string{}
	for _, v := range tickets {
		images[v[0]] = append(images[v[0]], v[1])
	}
	for k, _ := range images {
		sort.Strings(images[k])
	}
	path = []string{}
	findItineraryHelper("JFK")
	return path
}

func findItineraryHelper(start string) {
	for len(images[start]) != 0 {
		target := images[start][0]
		images[start] = images[start][1:]
		findItineraryHelper(target)
	}
	path = append(path, start)
}
