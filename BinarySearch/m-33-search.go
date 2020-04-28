package BinarySearch

/**
 * @Time: 2020/4/28 21:08
 * @DESC: 33. 搜索旋转排序数组
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 * 你可以假设数组中不存在重复的元素。
 * 你的算法时间复杂度必须是 O(log n) 级别。
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
 * @param nums
 * @param target
 * @return int
 */
func Search(nums []int, target int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return -1
	}
	if lenNums == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left := 0
	right := lenNums - 1
	mid := (left + right) / 2
	for left <= right {
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[mid] < target || nums[left] > target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] > target || nums[right] < target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		mid = (left + right) / 2
	}
	return -1
}
