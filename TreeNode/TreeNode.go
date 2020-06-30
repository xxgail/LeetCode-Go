package TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	k := arraySearch(preorder[0], inorder)
	root.Left = buildTree(preorder[1:k+1], inorder[0:k])
	root.Right = buildTree(preorder[k+1:], inorder[k+1:])
	return root
}

func arraySearch(i int, arr []int) int {
	for k, v := range arr {
		if v == i {
			return k
		}
	}
	return -1
}
