package main

import (
	"fmt"
	"github.com/xxgail/LeetCode-Go/BitManipulation"
)

func main() {
	//A := []int{12, -4, 16, -5, 9, -3, 3, 8, 0}
	//res := Array.CanThreePartsEqualSum(A)
	//fmt.Println(res)

	//str1 := "ABABAB"
	//str2 := "ABAB"
	//res := String.GocOfStrings(str1, str2)

	//nums := []int{1, 2, 1, 1}
	//res := Array.MajorityElement(nums)
	//fmt.Println(res)
	//[10,9,2,5,3,7,101,18]
	//[4,10,4,3,8,9]
	//nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//res := DynamicProgramming.LengthOfLIS(nums)
	//s := "catsandog"
	//wordDict := []string{"cats","dog","sand","and","cat"}
	//res := DynamicProgramming.WordBreak(s,wordDict)
	//res := Backtracking.GenerateParenthesis(3)
	//one := &LinkedList.ListNode{
	//	Val: 1,
	//	Next: &LinkedList.ListNode{
	//		Val: 3,
	//		Next: &LinkedList.ListNode{
	//			Val: 4,
	//		},
	//	},
	//}
	//two := &LinkedList.ListNode{
	//	Val: 1,
	//	Next: &LinkedList.ListNode{
	//		Val: 4,
	//		Next: &LinkedList.ListNode{
	//			Val: 5,
	//		},
	//	},
	//}
	//three := &LinkedList.ListNode{
	//	Val: 2,
	//	Next: &LinkedList.ListNode{
	//		Val: 6,
	//	},
	//}
	//lists := []*LinkedList.ListNode{one, two, three}
	//res := LinkedList.MergeKLists(lists)
	//fmt.Println(res)
	nums := []int{1, 2, 10, 3, 3, 4, 1, 4}
	res := BitManipulation.SingleNumbers(nums)
	fmt.Println(res)
}
