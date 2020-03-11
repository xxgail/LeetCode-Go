package main

import (
	"LeetCode-Go/Array"
	"fmt"
)

func main() {
	A := []int{12, -4, 16, -5, 9, -3, 3, 8, 0}
	res := Array.CanThreePartsEqualSum(A)
	fmt.Println(res)
}
