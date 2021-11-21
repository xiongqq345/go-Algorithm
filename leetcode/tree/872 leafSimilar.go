package tree

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var vals1, vals2 []int
	var preorder func(*TreeNode, *[]int)
	preorder = func(node *TreeNode, vals *[]int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			*vals = append(*vals, node.Val)
			return
		}
		preorder(node.Left, vals)
		preorder(node.Right, vals)
	}
	preorder(root1, &vals1)
	preorder(root2, &vals2)

	if len(vals1) != len(vals2) {
		return false
	}
	for i, v := range vals1 {
		if vals2[i] != v {
			return false
		}
	}
	return true
}
