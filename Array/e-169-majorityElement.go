package Array

/**
 * @Time: 2020/3/13 16:21
 * @DESC: 169. 多数元素 简单
 * 给定一个大小为 n 的数组，找到其中的多数元素。
 * 多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 * @params []int
 * @return int
 */
func MajorityElement(nums []int) int {
	//m := make(map[int]int)
	//var res int
	//length := len(nums)/2
	//for _,value := range nums{
	//	_, ok := m[value]
	//	if ok == true{
	//		m[value] = m[value] + 1
	//	}else{
	//		m[value] = 1
	//	}
	//	if m[value] > length{
	//		res = value
	//	}
	//}
	//return res

	var last, count int
	for _, v := range nums {
		if count == 0 {
			last = v
			count = 1
			continue
		}
		if last == v {
			count++
		} else {
			count--
		}
	}
	return last
}

// map类型不像其他基础类型，有初始值。所以在初始化的时候需要用make来初始化
// _,ok := map[k] 验证该集合中map[k]是否存在。若存在ok=true，否则ok=false
