package TreeNode

func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var item []*TreeNode
	nodeCount := 1
	nextNodeCount := 0
	item = append(item, root)
	var itemArr []int
	for len(item) != 0 {
		current := item[0]
		item = item[1:]
		itemArr = append(itemArr, current.Val)
		nodeCount--
		if current.Left != nil {
			item = append(item, current.Left)
			nextNodeCount++
		}
		if current.Right != nil {
			item = append(item, current.Right)
			nextNodeCount++
		}
		if nodeCount == 0 {
			res = append(res, itemArr)
			nodeCount = nextNodeCount
			nextNodeCount = 0
			itemArr = []int{}
		}
	}
	return res
}
