package BitManipulation

/**
 * @Time: 2020/04/28 21:37
 * @DESC: 面试题60-1. 数组中数字出现的次数
 * 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
 * 请写程序找出这两个只出现一次的数字。
 * 要求时间复杂度是O(n)，空间复杂度是O(1)。
 * https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
 * @param nums
 * @return array
 */
func SingleNumbers(nums []int) []int {
	res := 0

	for _, v := range nums {
		res ^= v
	}

	index := 1
	for res&index == 0 {
		index <<= 1
	}

	r1, r2 := 0, 0
	for _, v := range nums {
		if v&index == 0 {
			r1 ^= v
		} else {
			r2 ^= v
		}
	}

	return []int{r1, r2}
}
