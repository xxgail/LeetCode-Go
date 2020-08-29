package Heap

import (
	"container/heap"
	"fmt"
)

/**
 * @Date 2020/8/29
 * @Desc 378. 有序矩阵中第K小的元素
给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
示例：
	matrix = [
	   [ 1,  5,  9],
	   [10, 11, 13],
	   [12, 13, 15]
	],
	k = 8,
	返回 13。
 * @Param
 * @return
 * @link https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/
 **/
func KthSmallest(matrix [][]int, k int) int {
	h := &IHeap{}
	for i := 0; i < len(matrix); i++ {
		heap.Push(h, [3]int{matrix[i][0], i, 0})
	}
	fmt.Println(h)
	for i := 0; i < k-1; i++ {
		now := heap.Pop(h).([3]int)
		fmt.Println(h)
		if now[2] != len(matrix)-1 {
			heap.Push(h, [3]int{matrix[now[1]][now[2]+1], now[1], now[2] + 1})
		}
		fmt.Println(h)
	}
	return heap.Pop(h).([3]int)[0]
}

type IHeap [][3]int

func (h IHeap) Len() int {
	return len(h)
}
func (h IHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
