package main

import (
	"LeetCode-Go/DynamicProgramming"
	"fmt"
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
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := DynamicProgramming.LengthOfLIS(nums)
	fmt.Println(res)
}
