package LinkedList

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @Time: 2020/4/26 22:44
 * @DESC: 23. 合并K个排序链表
 * 链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
 * 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
 * 示例:
 * 输入:
 * [
 *   1->4->5,
 *   1->3->4,
 *   2->6
 * ]
 * 输出: 1->1->2->3->4->4->5->6
 * @params
 * @return
 */
func MergeKLists(lists []*ListNode) *ListNode {
	var arr []int
	// 转数组
	for _, v := range lists {
		for v != nil {
			arr = append(arr, v.Val)
			v = v.Next
		}
	}
	if len(arr) == 0 {
		return nil
	}
	sort.Ints(arr) // 排序
	res := &ListNode{
		Val: 1,
	}
	current := res
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{
			Val: arr[i],
		}
		current = current.Next
	}
	current.Next = nil
	return res
}
