package main

import (
	"LeetCode-Go/String"
	"fmt"
)

func main() {
	//A := []int{12, -4, 16, -5, 9, -3, 3, 8, 0}
	//res := Array.CanThreePartsEqualSum(A)
	//fmt.Println(res)

	str1 := "ABABAB"
	str2 := "ABAB"
	res := String.GocOfStrings(str1, str2)
	fmt.Println(res)
}
