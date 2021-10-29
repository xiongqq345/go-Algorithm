package bt

//给定一个二叉搜索树，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和。
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Right)
		sum += node.Val
		node.Val = sum
		inorder(node.Left)
	}
	inorder(root)
	return root
}
