package Array

/**
560. 和为K的子数组
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
示例 1 :
输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。

https://leetcode-cn.com/problems/subarray-sum-equals-k
*/
func SubarraySum(nums []int, k int) int {
	// hash表
	res := 0
	s := make(map[int]int)
	var sum int
	for _, n := range nums {
		sum += n
		if sum == k {
			res++
		}
		res += s[sum-k]
		s[sum]++
	}
	return res

	// 普通数组遍历
	//res := 0
	//numsLen := len(nums)
	//var sum int
	//for i := 0; i < numsLen; i++ {
	//	sum = 0
	//	for j := i; j >= 0; j-- {
	//		sum += nums[j]
	//		if sum == k {
	//			res++
	//		}
	//	}
	//}
	//return res
}
