package DynamicProgramming

// https://leetcode-cn.com/problems/minimum-cost-for-tickets/

func MinCostTickets(days []int, costs []int) int {
	if len(days) == 0 {
		return 0
	}

	res := make(map[int]int)
	min := costs[0]
	if min > costs[1] {
		min = costs[1]
	}
	if min > costs[2] {
		min = costs[2]
	}
	res[days[0]] = min
	lastDay := days[len(days)-1]

	for i := days[0] + 1; i <= lastDay; i++ {
		ok := false
		for _, v := range days {
			if i == v {
				ok = true
				break
			}
		}
		if ok == true {
			var i1 = 0
			if i-1 >= days[0] {
				i1 = res[i-1]
			}
			var i7 = 0
			if i-7 >= days[0] {
				i7 = res[i-7]
			}
			var i30 = 0
			if i-30 >= days[0] {
				i30 = res[i-30]
			}
			if i1+costs[0] < i7+costs[1] {
				min = i1 + costs[0]
			} else {
				min = i7 + costs[1]
			}
			if min > i30+costs[2] {
				min = i30 + costs[2]
			}
		}
		res[i] = min
	}
	return res[lastDay]
}
