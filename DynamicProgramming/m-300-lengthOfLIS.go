package DynamicProgramming

/**
 * @Time: 2020/3/14 22:42
 * @DESC: 300. 最长上升子序列 中等
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *
 * 示例:
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 *
 * 说明: 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * @param []int
 * @return int
 */
func LengthOfLIS(nums []int) int {
	// 动态规划
	if len(nums) <= 1 {
		return len(nums)
	}
	//dp := make(map[int]int)
	//dp[0] = 1
	//res := 1
	//for i := 1; i < len(nums); i++ {
	//	max := 0
	//	for j := 0; j < i; j++ {
	//		if nums[i] > nums[j]{
	//			if max < dp[j] {
	//				max = dp[j]
	//			}
	//		}
	//	}
	//	dp[i] = max + 1
	//	if res < dp[i] {
	//		res = dp[i]
	//	}
	//}
	//return res

	var res []int
	res = append(res, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] > res[len(res)-1] {
			res = append(res, nums[i])
		} else {
			for j := 0; j < len(res); j++ {
				if nums[i] <= res[j] {
					res[j] = nums[i]
					break
				}
			}
		}
	}
	return len(res)
}
