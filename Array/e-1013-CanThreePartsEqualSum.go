package Array

/**
 * @Time: 2020/3/11 18:08
 * @DESC: 1013. 将数组分成和相等的三个部分 简单
 * 给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false
 * 形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。
 * @param A
 * @return bool
 */
func CanThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, value := range A {
		sum += value
	}
	if sum%3 != 0 {
		return false
	}
	avg := sum / 3

	s := 0
	count := 0
	for _, value := range A {
		s += value
		if s == avg {
			s = 0
			count++
		}
	}
	if count >= 3 {
		return true
	} else {
		return false
	}
}
